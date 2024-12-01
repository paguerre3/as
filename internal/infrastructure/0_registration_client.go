package infrastructure

import comm "github.com/paguerre3/as/internal/common"

func (c *clientHandlerImpl) Register(alias, country, email, applyRole string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"alias":      alias,
		"country":    country,
		"email":      email,
		"apply_role": applyRole, // engineering
	}

	uri := comm.BuildASApiUri(1, "register")

	// Send the POST request
	resp, err := c.client.R().
		SetHeader(comm.CONTENT_TYPE, comm.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return handleError(resp, err)
	}

	return handleResponse(resp)
}
