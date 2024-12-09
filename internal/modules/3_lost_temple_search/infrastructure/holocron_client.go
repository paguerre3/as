package infrastructure

import (
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type HolocronClient interface {
	Fetch(uri string) (map[string]interface{}, int, error)
	FetchSWAPIPlanets(index int) (map[string]interface{}, int, error)
	QueryOracle(name string) (map[string]interface{}, int, error)
	OracleSolution(balancedBlanet string) (map[string]interface{}, int, error)
}

type holocronClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewHolocronClient() HolocronClient {
	return &holocronClientImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *holocronClientImpl) Fetch(uri string) (map[string]interface{}, int, error) {
	// Send the GET request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		Get(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}

func (c *holocronClientImpl) fetchSWAPI(path string, index int) (map[string]interface{}, int, error) {
	uri := common_infra.BuilSWAPIPeopleUri(path, index)
	return c.Fetch(uri)
}

func (c *holocronClientImpl) FetchSWAPIPlanets(index int) (map[string]interface{}, int, error) {
	return c.fetchSWAPI("planets", index)
}

func (c *holocronClientImpl) QueryOracle(name string) (map[string]interface{}, int, error) {
	uri := common_infra.BuildASApiUri(1, "s1/e3/resources/oracle-rolodex")

	// Send the GET request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetQueryParam("name", name).
		Get(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}

	return c.clientHandler.HandleResponse(resp)
}

func (c *holocronClientImpl) OracleSolution(balancedBlanet string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"planet": balancedBlanet,
	}

	uri := common_infra.BuildASApiUri(1, "s1/e3/solution")

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
