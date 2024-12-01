package infrastructure

func (c *clientHandlerImpl) StartBattle() (string, int, error) {
	uri := buildASApiUri(1, "s1/e5/actions/start")
	resp, err := c.client.R().
		SetHeader(AUTHORIZATION, BEARER_API_KEY).
		Post(uri)
	if err != nil {
		return handleStringError(resp, err)
	}
	return handleStringResponse(resp)
}

func (c *clientHandlerImpl) PerformTurn(action string, x string, y int) (map[string]interface{}, int, error) {
	requestBody := map[string]interface{}{
		"action": action,
		"attack_position": map[string]interface{}{
			"x": x,
			"y": y,
		},
	}
	uri := buildASApiUri(1, "s1/e5/actions/perform-turn")
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
