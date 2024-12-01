package infrastructure

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
