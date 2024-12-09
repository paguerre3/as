package infrastructure

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

type ClientHandler interface {
	HandleResponse(resp *resty.Response) (map[string]interface{}, int, error)
	HandleError(resp *resty.Response, err error) (map[string]interface{}, int, error)
	HandleArrayResponse(resp *resty.Response) ([]map[string]interface{}, int, error)
	HandleArrayError(resp *resty.Response, err error) ([]map[string]interface{}, int, error)
	HandleStringResponse(resp *resty.Response) (string, int, error)
	HandleStringError(resp *resty.Response, err error) (string, int, error)
	Client() *resty.Client
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

func handleResponse[T any](resp *resty.Response, debug bool) (T, int, error) {
	var response T
	if debug {
		logResponse(resp)
	}
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		var zero T
		return zero, resp.StatusCode(), err
	}
	return response, resp.StatusCode(), nil
}

func handleError[T any](resp *resty.Response, err error) (T, int, error) {
	var statusCode int
	if resp != nil {
		statusCode = resp.StatusCode()
	}
	var zero T
	return zero, statusCode, err
}

func (c *clientHandlerImpl) HandleResponse(resp *resty.Response) (map[string]interface{}, int, error) {
	return handleResponse[map[string]interface{}](resp, c.debug)
}

func (c *clientHandlerImpl) HandleArrayResponse(resp *resty.Response) ([]map[string]interface{}, int, error) {
	return handleResponse[[]map[string]interface{}](resp, c.debug)
}

func (c *clientHandlerImpl) HandleStringResponse(resp *resty.Response) (response string, statusCode int, err error) {
	return handleResponse[string](resp, c.debug)
}

func (c *clientHandlerImpl) HandleError(resp *resty.Response, err error) (map[string]interface{}, int, error) {
	return handleError[map[string]interface{}](resp, err)
}

func (c *clientHandlerImpl) HandleArrayError(resp *resty.Response, err error) ([]map[string]interface{}, int, error) {
	return handleError[[]map[string]interface{}](resp, err)
}

func (c *clientHandlerImpl) HandleStringError(resp *resty.Response, err error) (string, int, error) {
	return handleError[string](resp, err)
}

func (c *clientHandlerImpl) Client() *resty.Client {
	return c.client
}

func VerifyCorrectness(t *testing.T, response map[string]interface{}, statusCode int, err error) {
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Printf("response: %+v", response)

	resultValue := response["result"]
	assert.NotEmpty(t, resultValue)
	assert.Equal(t, "correct", resultValue)
}
