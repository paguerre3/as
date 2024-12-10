package application

import (
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/modules/3_lost_temple_search/domain"
)

type SearchLostTempleUseCase interface {
	Execute() (map[string]interface{}, int, error)
}

type HolocronClient interface {
	Fetch(uri string) (map[string]interface{}, int, error)
	FetchSWAPIPlanets(index int) (map[string]interface{}, int, error)
	QueryOracle(name string) (map[string]interface{}, int, error)
	OracleSolution(balancedBlanet string) (map[string]interface{}, int, error)
} // Exposing duplicated interface to avoid DDD violations

type searchLostTempleImpl struct {
	holocronClient HolocronClient
}

func NewSearchLostTempleUseCase(holocronClient HolocronClient) SearchLostTempleUseCase {
	return &searchLostTempleImpl{
		holocronClient: holocronClient,
	}
}

func (s *searchLostTempleImpl) Execute() (map[string]interface{}, int, error) {
	// Actual test logic goes here:
	planets, err := fetchAllPlanets(s.holocronClient)
	if err != nil {
		return nil, 0, err
	}
	if len(planets) == 0 {
		return nil, 0, fmt.Errorf("no planets found")
	}
	for _, planet := range planets {
		ibf, err := calculateIBF(s.holocronClient, planet)
		if err != nil {
			// only possible error is "no residents found" whihc produces 0 IBF
			log.Warnf("Error calculating IBF for planet %s: %v", planet.Name, err)
		}
		if ibf == 0 && err == nil {
			// only one panet with people and balanced (IBF = 0)
			response, statusCode, err := s.holocronClient.OracleSolution(planet.Name)
			log.Infof(strings.Repeat("-", 75))
			log.Infof(strings.Repeat("-", 75))
			// "Balanced Planet: Ryloth"
			log.Infof("Balanced Planet: %s", planet.Name)
			log.Infof(strings.Repeat("-", 75))
			return response, statusCode, err
		}
	}
	return nil, 0, nil
}

func fetchAllPlanets(handler HolocronClient) ([]domain.Planet, error) {
	planets := []domain.Planet{}
	// fech planets sequentially as a safe approach:
	for i := 1; ; i++ {
		resp, statusCode, err := handler.FetchSWAPIPlanets(i)
		if statusCode == 404 {
			// planet retrieval ended
			break
		}
		if err != nil {
			return nil, err
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("failed to fetch SWAPI planets: status code %d", statusCode)
		}

		name, ok := resp["name"]
		if !ok {
			return nil, fmt.Errorf("failed to fetch SWAPI planets: missing name")
		}
		residents, ok := resp["residents"]
		if !ok {
			return nil, fmt.Errorf("failed to fetch SWAPI planets: missing residents")
		}

		ress := residents.([]interface{})
		var residentNames []string
		for _, resident := range ress {
			if name, ok := resident.(string); ok {
				residentNames = append(residentNames, name)
			}
		}

		// only 2 fields being used for best performance:
		planets = append(planets, domain.Planet{
			Name:      name.(string),
			Residents: residentNames,
		})
	}
	return planets, nil
}

func calculateIBF(handler HolocronClient, planet domain.Planet) (float64, error) {
	var lightSideCount, darkSideCount int

	var wg sync.WaitGroup
	var mu sync.Mutex

	// for each resident/by gorutine, query to check dark or light side:
	for _, residentURL := range planet.Residents {
		wg.Add(1)
		go func(residentURL string) {
			defer wg.Done()

			// Obtain people information from planet
			resp, statusCode, err := handler.Fetch(residentURL)
			if err != nil {
				log.Errorf("Error fetching person: %v", err)
				return
			}
			if statusCode != 200 {
				log.Errorf("Error fetching person: status code %d", statusCode)
				return
			}

			name, ok := resp["name"]
			if !ok {
				log.Errorf("Error fetching person: missing name")
				return
			}

			// Query light or dark side:
			side, statusCode, err := handler.QueryOracle(name.(string))
			if err != nil {
				log.Printf("Error querying oracle: %v", err)
				return
			}
			if statusCode != 200 {
				log.Printf("Error querying oracle: status code %d", statusCode)
				return
			}

			oracleNotes, ok := side["oracle_notes"]
			if !ok {
				log.Printf("Error querying oracle: missing oracle_notes")
				return
			}
			//log.Infof("Oracle notes for %s: %s", name, oracleNotes)

			decoded, err := base64.StdEncoding.DecodeString(oracleNotes.(string))
			if err != nil {
				log.Errorf("Error decoding oracle notes: %v", err)
			}

			// protect concurrent count
			mu.Lock()
			notes := strings.ToLower(string(decoded))
			if strings.Contains(notes, "dark") {
				darkSideCount++
			} else if strings.Contains(notes, "light") {
				lightSideCount++
			}
			mu.Unlock()
		}(residentURL)
	}

	// wait for all goroutines to end
	wg.Wait()

	return domain.CalculateIBF(lightSideCount, darkSideCount)
}
