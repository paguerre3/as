package infrastructure

import (
	"encoding/json"
	"fmt"
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

func (c *clientHandlerImpl) Register(alias, country, email, applyRole string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"alias":      alias,
		"country":    country,
		"email":      email,
		"apply_role": applyRole, // engineering
	}

	uri := buildASApiUri(1, "register")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return handleResponse(resp)
}

func (c *clientHandlerImpl) Measurement() (map[string]interface{}, int, error) {
	uri := buildASApiUri(1, "s1/e1/resources/measurement")

	// Send the GET request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		Get(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return handleResponse(resp)
}

func (c *clientHandlerImpl) MeassurmentSolution(speed int) (map[string]interface{}, int, error) {
	requestBody := map[string]int{
		"speed": speed,
	}

	uri := buildASApiUri(1, "s1/e1/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return handleResponse(resp)
}

func (c *clientHandlerImpl) FetchStars(page int) ([]map[string]interface{}, int, error) {
	uri := buildASApiUri(1, "s1/e2/resources/stars")

	// Send the GET request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetQueryParam("page", fmt.Sprintf("%d", page)).
		Get(uri)
	if err != nil {
		return handleArrayError(resp, err)
	}

	return handleArrayResponse(resp)
}

func (c *clientHandlerImpl) ResonanceSolution(averageResonance int) (map[string]interface{}, int, error) {
	requestBody := map[string]int{
		"average_resonance": averageResonance,
	}

	uri := buildASApiUri(1, "s1/e2/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return handleResponse(resp)
}

func (c *clientHandlerImpl) Fetch(uri string) (map[string]interface{}, int, error) {
	// Send the GET request
	resp, err := c.client.R().
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		Get(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

func (c *clientHandlerImpl) fetchSWAPI(path string, index int) (map[string]interface{}, int, error) {
	uri := builSWAPIPeopleUri(path, index)
	return c.Fetch(uri)
}

func (c *clientHandlerImpl) FetchSWAPIPlanets(index int) (map[string]interface{}, int, error) {
	return c.fetchSWAPI("planets", index)
}

func (c *clientHandlerImpl) QueryOracle(name string) (map[string]interface{}, int, error) {
	uri := buildASApiUri(1, "s1/e3/resources/oracle-rolodex")

	// Send the GET request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetQueryParam("name", name).
		Get(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return handleResponse(resp)
}

func (c *clientHandlerImpl) OracleSolution(balancedBlanet string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"planet": balancedBlanet,
	}

	uri := buildASApiUri(1, "s1/e3/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

func (c *clientHandlerImpl) UserAndPasswordSolution(username, password string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"username": username,
		"password": password,
	}
	uri := buildASApiUri(1, "s1/e4/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

func (c *clientHandlerImpl) StartBattle() (string, int, error) {
	uri := buildASApiUri(1, "s1/e5/actions/start")
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		Post(uri)
	if err != nil {
		return handleStringError(resp, err)
	}
	return handleStringResponse(resp)
}

func (c *clientHandlerImpl) PerformTurn(action string, x string, y int) (map[string]interface{}, int, error) {
	requestBody := map[string]interface{}{
		"action": action,
		"attack_position": map[string]interface{}{
			"x": x,
			"y": y,
		},
	}
	uri := buildASApiUri(1, "s1/e5/actions/perform-turn")
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		SetHeader(CONTENT_TYPE, APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

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