package infrastructure

import "fmt"

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
