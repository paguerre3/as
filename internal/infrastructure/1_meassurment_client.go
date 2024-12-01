package infrastructure

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
