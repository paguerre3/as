package infrastructure

import (
	"encoding/json"
	"fmt"
	"sync"
)

func (c *clientHandlerImpl) GetPockemonTypes() (map[string]interface{}, int, error) {
	uri := buildPockeApi("type")
	resp, err := c.client.R().
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		Get(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

func (c *clientHandlerImpl) GetTypeData(typeUrl, typeName string) (map[string]interface{}, int, error) {
	resp, err := c.client.R().
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		Get(typeUrl)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

func (c *clientHandlerImpl) GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error) {
	resp, err := c.client.R().
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		Get(pokemonUrl)
	if resp != nil {
		statusCode = resp.StatusCode()
	}
	if err != nil {
		return statusCode, fmt.Errorf("error retrieving pokemon: %v", err)
	}
	var pokemon map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &pokemon); err != nil {
		return statusCode, fmt.Errorf("error unmarshaling pokemon: %v", err)
	}
	height, ok := pokemon["height"].(float64)
	if !ok {
		return statusCode, fmt.Errorf("error obtaining Pokemon height: %v", pokemon["height"])
	}
	mu.Lock()
	typeHeights[typeName] = append(typeHeights[typeName], height)
	mu.Unlock()
	return statusCode, nil
}

func (c *clientHandlerImpl) PokemonSolution(pokeSolution map[string]interface{}) (map[string]interface{}, int, error) {
	uri := buildASApiUri(1, "s1/e6/solution")
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(pokeSolution).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}
