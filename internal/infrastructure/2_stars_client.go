package infrastructure

import (
	"fmt"

	comm "github.com/paguerre3/as/internal/common"
)

func (c *clientHandlerImpl) FetchStars(page int) ([]map[string]interface{}, int, error) {
	uri := comm.BuildASApiUri(1, "s1/e2/resources/stars")

	// Send the GET request
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
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

	uri := comm.BuildASApiUri(1, "s1/e2/solution")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(comm.AUTHORIZATION, comm.BEARER_API_KEY).
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return c.handleResponse(resp)
}
