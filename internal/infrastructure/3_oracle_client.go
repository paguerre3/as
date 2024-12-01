package infrastructure

import (
	comm "github.com/paguerre3/as/internal/common"
)

func (c *clientHandlerImpl) Fetch(uri string) (map[string]interface{}, int, error) {
	// Send the GET request
	resp, err := c.client.R().
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		Get(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}

func (c *clientHandlerImpl) fetchSWAPI(path string, index int) (map[string]interface{}, int, error) {
	uri := comm.BuilSWAPIPeopleUri(path, index)
	return c.Fetch(uri)
}

func (c *clientHandlerImpl) FetchSWAPIPlanets(index int) (map[string]interface{}, int, error) {
	return c.fetchSWAPI("planets", index)
}

func (c *clientHandlerImpl) QueryOracle(name string) (map[string]interface{}, int, error) {
	uri := comm.BuildASApiUri(1, "s1/e3/resources/oracle-rolodex")

	// Send the GET request
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
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

	uri := comm.BuildASApiUri(1, "s1/e3/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}
	return handleResponse(resp)
}
