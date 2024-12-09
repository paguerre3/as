package application

import (
	"fmt"
	"sort"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/modules/6_prism_city_infiltration/domain"
	"github.com/shopspring/decimal"
)

type CalculatePokemonTypesAverageHeightsUseCase interface {
	Execute() (map[string]interface{}, int, error)
}

type PokemonClient interface {
	GetPockemonTypes() (map[string]interface{}, int, error)
	GetTypeData(typeUrl, typeName string) (map[string]interface{}, int, error)
	GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error)
	PokemonSolution(pokeSolution map[string]interface{}) (map[string]interface{}, int, error)
} // Exposing duplicated interface to avoid DDD violations

type calculatePokemonTypesAverageHeightsUseCaseImpl struct {
	pokemonClient PokemonClient
}

func NewCalculatePokemonTypesAverageHeightsUseCase(pokemonClient PokemonClient) CalculatePokemonTypesAverageHeightsUseCase {
	return &calculatePokemonTypesAverageHeightsUseCaseImpl{
		pokemonClient: pokemonClient,
	}
}

func (c *calculatePokemonTypesAverageHeightsUseCaseImpl) Execute() (map[string]interface{}, int, error) {
	typesResponse, statusCode, err := c.pokemonClient.GetPockemonTypes()
	if err != nil {
		return nil, 0, fmt.Errorf("error getting pokemon types: %v", err)
	}
	if statusCode != 200 {
		return nil, 0, fmt.Errorf("error getting pokemon types: status code %d", statusCode)
	}
	typesResults, ok := typesResponse["results"]
	if !ok {
		return nil, 0, fmt.Errorf("error getting pokemon types: missing results")
	}

	typeHeights := make(map[string][]float64)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, typeResult := range typesResults.([]interface{}) {
		wg.Add(1)
		go func(typeResult interface{}) {
			defer wg.Done()

			typeInfo, ok := typeResult.(map[string]interface{})
			if !ok {
				log.Errorf("error getting pokemon types: missing type info")
				return
			}
			pokemonType, ok := typeInfo["name"]
			if !ok {
				log.Errorf("error getting pokemon types: missing type name")
				return
			}
			typeURL, ok := typeInfo["url"]
			if !ok {
				log.Errorf("error getting pokemon types: missing type url")
				return
			}
			singleTypeResponse, statusCode, err := c.pokemonClient.GetTypeData(typeURL.(string), pokemonType.(string))
			if err != nil {
				log.Errorf("error getting signle pokemon type: %v", err)
				return
			}
			if statusCode != 200 {
				log.Errorf("error getting signle pokemon type: status code %d", statusCode)
				return
			}
			pokemonsByType, ok := singleTypeResponse["pokemon"]
			if !ok {
				log.Errorf("error getting pokemon by type %s: missing pokemon results", pokemonType.(string))
				return
			}

			for _, pokemon := range pokemonsByType.([]interface{}) {
				pokemonContainer, ok := pokemon.(map[string]interface{})
				if !ok {
					log.Errorf("error getting pokemon by type %s: missing pokemon info", pokemonType.(string))
					return
				}
				pokemonInfo, ok := pokemonContainer["pokemon"]
				if !ok {
					log.Errorf("error getting pokemon by type %s: missing pokemon name", pokemonType.(string))
					return
				}
				pokemonName, ok := pokemonInfo.(map[string]interface{})["name"]
				if !ok {
					log.Errorf("error getting pokemon name by type %s: missing pokemon name", pokemonType.(string))
					return
				}
				pokemonUrl, ok := pokemonInfo.(map[string]interface{})["url"]
				if !ok {
					log.Errorf("error getting pokemon url by type %s with name %s: missing pokemon url", pokemonType.(string), pokemonName.(string))
					return
				}
				statusCode, err := c.pokemonClient.GetUpdatePokemonHeight(pokemonUrl.(string), pokemonType.(string), typeHeights, &mu)
				if err != nil {
					log.Errorf("error getting pokemon height by url %s with name %s: %v", pokemonUrl.(string), pokemonName.(string), err)
					return
				}
				if statusCode != 200 {
					log.Errorf("error getting pokemon height by url %s with name %s: status code %d", pokemonUrl.(string), pokemonName.(string), statusCode)
					return
				}
			}
		}(typeResult)
	}
	wg.Wait()

	averageHeights := domain.CalculateAverageHeights(typeHeights)

	keys := make([]string, 0, len(averageHeights))
	for key := range averageHeights {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	orderedMap := make(map[string]float64)
	for _, key := range keys {
		avgh, _ := decimal.NewFromFloat(averageHeights[key]).Round(3).Float64()
		orderedMap[key] = avgh
	}
	solution := map[string]interface{}{
		"heights": orderedMap,
	}
	log.Infof("solution: %s\n", solution)
	return c.pokemonClient.PokemonSolution(solution)
}
