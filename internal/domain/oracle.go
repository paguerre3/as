package domain

import (
	"encoding/base64"
	"fmt"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
)

type Planet struct {
	Name      string   `json:"name"`
	Residents []string `json:"residents"`
}

func AllPlanets(handler ClientHandler) ([]Planet, error) {
	planets := []Planet{}
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
		planets = append(planets, Planet{
			Name:      name.(string),
			Residents: residentNames,
		})
	}
	return planets, nil
}

func CalculateIBF(handler ClientHandler, planet Planet) (float64, error) {
	var lightSideCount, darkSideCount int

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Para cada residente, consultar su lado en la Fuerza
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

			// Contar en base al lado
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

	// Esperar a que todas las goroutines terminen
	wg.Wait()

	// Calcular el IBF
	totalCount := lightSideCount + darkSideCount
	if totalCount == 0 {
		return 0, fmt.Errorf("no residents found")
	}
	ibf := float64(lightSideCount-darkSideCount) / float64(totalCount)
	return ibf, nil
}
