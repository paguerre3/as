package infrastructure

import (
	"encoding/json"
	"fmt"
	"sync"

	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type PokemonClient interface {
	GetPockemonTypes() (map[string]interface{}, int, error)
	GetTypeData(typeUrl, typeName string) (map[string]interface{}, int, error)
	GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error)
	PokemonSolution(pokeSolution map[string]interface{}) (map[string]interface{}, int, error)
}

type pokemonClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewPokemonClient() PokemonClient {
	return &pokemonClientImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *pokemonClientImpl) GetPockemonTypes() (map[string]interface{}, int, error) {
	uri := common_infra.BuildPockeApi("type")
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		Get(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}

func (c *pokemonClientImpl) GetTypeData(typeUrl, typeName string) (map[string]interface{}, int, error) {
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		Get(typeUrl)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}

func (c *pokemonClientImpl) GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error) {
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
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

func (c *pokemonClientImpl) PokemonSolution(pokeSolution map[string]interface{}) (map[string]interface{}, int, error) {
	uri := common_infra.BuildASApiUri(1, "s1/e6/solution")
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetBody(pokeSolution).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}
