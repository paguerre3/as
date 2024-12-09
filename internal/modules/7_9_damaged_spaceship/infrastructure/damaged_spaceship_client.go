package infrastructure

import (
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

const (
	e7 = "s1/e7/solution"
	e9 = "s1/e9/solution"
)

type DamagedSpaceshipClient interface {
	RegisterEndpont7Solution(endpoint string) (map[string]interface{}, int, error)
	RegisterEndpont9Solution(endpoint string) (map[string]interface{}, int, error)
}

type damagedSpaceshipClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewDamagedSpaceshipClient() DamagedSpaceshipClient {
	return &damagedSpaceshipClientImpl{
		clientHandler: common_infra.NewClientHandlerDebug(),
	}
}

func (c *damagedSpaceshipClientImpl) registerEndpontSolution(endpoint, solution string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"base_url": endpoint,
	}
	uri := common_infra.BuildASApiUri(1, solution)

	// Send the POST request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}

func (c *damagedSpaceshipClientImpl) RegisterEndpont7Solution(endpoint string) (map[string]interface{}, int, error) {
	return c.registerEndpontSolution(endpoint, e7)
}

func (c *damagedSpaceshipClientImpl) RegisterEndpont9Solution(endpoint string) (map[string]interface{}, int, error) {
	return c.registerEndpontSolution(endpoint, e9)
}
