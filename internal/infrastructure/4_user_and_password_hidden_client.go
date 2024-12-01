package infrastructure

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
