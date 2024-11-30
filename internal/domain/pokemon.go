package domain

import (
	"fmt"
	"sort"
	"sync"

	"github.com/labstack/gommon/log"
)

type TypeResponse struct {
	Results []TypeInfo `json:"results"`
}

type TypeInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Height int `json:"height"`
}

type Solution struct {
	Heights map[string]float64 `json:"heights"`
}

func calculateAverageHeights(typeHeights map[string][]float64) map[string]float64 {
	averageHeights := make(map[string]float64)
	for typeName, heights := range typeHeights {
		var total float64
		for _, height := range heights {
			total += height
		}
		averageHeights[typeName] = float64(total) / float64(len(heights))
	}
	return averageHeights
}

func CalculatePokemonTypesAverageHeights(handler ClientHandler) (map[string]interface{}, error) {
	typesResponse, statusCode, err := handler.GetPockemonTypes()
	if err != nil {
		return nil, fmt.Errorf("error getting pokemon types: %v", err)
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("error getting pokemon types: status code %d", statusCode)
	}
	typesResults, ok := typesResponse["results"]
	if !ok {
		return nil, fmt.Errorf("error getting pokemon types: missing results")
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
			singleTypeResponse, statusCode, err := handler.GetTypeData(typeURL.(string), pokemonType.(string))
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
				statusCode, err := handler.GetUpdatePokemonHeight(pokemonUrl.(string), pokemonType.(string), typeHeights, &mu)
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

	averageHeights := calculateAverageHeights(typeHeights)
	sortedTypeNames := make([]string, 0, len(averageHeights))
	for typeName := range averageHeights {
		sortedTypeNames = append(sortedTypeNames, typeName)
	}
	sort.Strings(sortedTypeNames)

	heights := make(map[string]float64)
	for _, typeName := range sortedTypeNames {
		heights[typeName] = averageHeights[typeName]
	}
	solution := map[string]interface{}{
		"heights": heights,
	}
	return solution, nil
}
