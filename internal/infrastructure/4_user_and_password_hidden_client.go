package infrastructure

import comm "github.com/paguerre3/as/internal/common"

func (c *clientHandlerImpl) UserAndPasswordSolution(username, password string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"username": username,
		"password": password,
	}
	uri := comm.BuildASApiUri(1, "s1/e4/solution")

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
