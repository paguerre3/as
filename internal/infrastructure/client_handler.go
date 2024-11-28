package infrastructure

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ClientHandler interface {
	Register(alias, country, email, applyRole string) (response map[string]interface{}, statusCode int, err error)

	Measurement() (response map[string]interface{}, statusCode int, err error)
	MeassurmentSolution(speed int) (response map[string]interface{}, statusCode int, err error)

	FetchStars(page int) (response []map[string]interface{}, statusCode int, err error)
	ResonanceSolution(averaegeResonance int) (response map[string]interface{}, statusCode int, err error)
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

	uri := buildApiUri(1, "register")

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
	uri := buildApiUri(1, "s1/e1/resources/measurement")

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

	uri := buildApiUri(1, "s1/e1/solution")

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
	uri := buildApiUri(1, "s1/e2/resources/stars")

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

	uri := buildApiUri(1, "s1/e2/solution")

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
