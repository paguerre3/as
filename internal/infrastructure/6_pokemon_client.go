package infrastructure

import (
	"encoding/json"
	"fmt"
	"sync"

	comm "github.com/paguerre3/as/internal/common"
)

func (c *clientHandlerImpl) GetPockemonTypes() (map[string]interface{}, int, error) {
	uri := comm.BuildPockeApi("type")
	resp, err := c.client.R().
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		Get(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return c.handleResponse(resp)
}

func (c *clientHandlerImpl) GetTypeData(typeUrl, typeName string) (map[string]interface{}, int, error) {
	resp, err := c.client.R().
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		Get(typeUrl)
	if err != nil {
		return handleError(resp, err)
	}
	return c.handleResponse(resp)
}

func (c *clientHandlerImpl) GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error) {
	resp, err := c.client.R().
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
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
	uri := comm.BuildASApiUri(1, "s1/e6/solution")
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		SetBody(pokeSolution).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return c.handleResponse(resp)
}
