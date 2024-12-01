package infrastructure

import (
	"encoding/json"
	"sync"

	"github.com/go-resty/resty/v2"
)

type ClientHandler interface {
	Register(alias, country, email, applyRole string) (response map[string]interface{}, statusCode int, err error)

	Measurement() (response map[string]interface{}, statusCode int, err error)
	MeassurmentSolution(speed int) (response map[string]interface{}, statusCode int, err error)

	FetchStars(page int) (response []map[string]interface{}, statusCode int, err error)
	ResonanceSolution(averaegeResonance int) (response map[string]interface{}, statusCode int, err error)

	Fetch(uri string) (response map[string]interface{}, statusCode int, err error)
	FetchSWAPIPlanets(index int) (response map[string]interface{}, statusCode int, err error)
	QueryOracle(name string) (response map[string]interface{}, statusCode int, err error)
	OracleSolution(balancedBlanet string) (response map[string]interface{}, statusCode int, err error)

	UserAndPasswordSolution(username, password string) (response map[string]interface{}, statusCode int, err error)

	StartBattle() (response string, statusCode int, err error)
	PerformTurn(action string, x string, y int) (response map[string]interface{}, statusCode int, err error)

	GetPockemonTypes() (response map[string]interface{}, statusCode int, err error)
	GetTypeData(typeUrl, typeName string) (response map[string]interface{}, statusCode int, err error)
	GetUpdatePokemonHeight(pokemonUrl, typeName string, typeHeights map[string][]float64, mu *sync.Mutex) (statusCode int, err error)
	PokemonSolution(pokeSolution map[string]interface{}) (response map[string]interface{}, statusCode int, err error)
}

type clientHandlerImpl struct {
	client *resty.Client
}

func NewClientHandler() ClientHandler {
	return &clientHandlerImpl{
		client: resty.New(),
	}
}

func handleResponse(resp *resty.Response) (map[string]interface{}, int, error) {
	var response map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, resp.StatusCode(), err
	}
	return response, resp.StatusCode(), nil
}

func handleArrayResponse(resp *resty.Response) ([]map[string]interface{}, int, error) {
	var response []map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return nil, resp.StatusCode(), err
	}
	return response, resp.StatusCode(), nil
}

func handleStringResponse(resp *resty.Response) (response string, statusCode int, err error) {
	statusCode = resp.StatusCode()
	if err = json.Unmarshal(resp.Body(), &response); err != nil {
		return "", statusCode, err
	}
	return response, statusCode, nil
}

func handleError(resp *resty.Response, err error) (map[string]interface{}, int, error) {
	var statusCode int
	if resp != nil {
		statusCode = resp.StatusCode()
	}
	return nil, statusCode, err
}

func handleArrayError(resp *resty.Response, err error) ([]map[string]interface{}, int, error) {
	var statusCode int
	if resp != nil {
		statusCode = resp.StatusCode()
	}
	return nil, statusCode, err
}

func handleStringError(resp *resty.Response, err error) (string, int, error) {
	var statusCode int
	if resp != nil {
		statusCode = resp.StatusCode()
	}
	return "", statusCode, err
}
