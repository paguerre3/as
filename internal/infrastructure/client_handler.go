package infrastructure

import (
	"encoding/json"
	"log"
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

	OpenDoor(body interface{}, gryffindorCookies *[]string) (response map[string]interface{}, statusCode int, err error)
	FirstClues() (response map[string]interface{}, statusCode int, err error)
	FourthClue(gryffindorCookies *[]string) (response map[string]interface{}, statusCode int, err error)
	HiddenMessageSolution(hiddenMessagePayload map[string]interface{}) (map[string]interface{}, int, error)

	RegisterEndpont9Solution(endpoint string) (map[string]interface{}, int, error)
}

type clientHandlerImpl struct {
	client *resty.Client
	debug  bool
}

func NewClientHandler() ClientHandler {
	return &clientHandlerImpl{
		client: resty.New(),
	}
}

func NewClientHandlerDebug() ClientHandler {
	return &clientHandlerImpl{
		client: resty.New(),
		debug:  true,
	}
}

func logResponse(resp *resty.Response) {
	if resp == nil {
		log.Println("Response is nil")
		return
	}

	// Log Status Code
	log.Printf("Status: %d", resp.StatusCode())

	// Log Headers
	log.Println("Headers:")
	for key, values := range resp.Header() {
		log.Printf("%s: %s", key, values)
	}

	// Log Raw Body
	log.Println("Body:")
	body := resp.String()
	if len(body) > 0 {
		log.Println(body)
	} else {
		log.Println("No body content")
	}

	// Log Full Raw Response
	log.Println("Raw Response Details:")
	log.Printf("%+v\n", resp.RawResponse)
}

func (c *clientHandlerImpl) handleResponse(resp *resty.Response) (map[string]interface{}, int, error) {
	var response map[string]interface{}
	if c.debug {
		logResponse(resp)
	}

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
