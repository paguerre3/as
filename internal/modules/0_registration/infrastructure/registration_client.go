package infrastructure

import (
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type RegistrationClient interface {
	Register(alias, country, email, applyRole string) (response map[string]interface{}, statusCode int, err error)
}

type registrationClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewRegistrationClient() RegistrationClient {
	return &registrationClientImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *registrationClientImpl) Register(alias, country, email, applyRole string) (response map[string]interface{}, statusCode int, err error) {
	requestBody := map[string]string{
		"alias":      alias,
		"country":    country,
		"email":      email,
		"apply_role": applyRole, // engineering
	}

	uri := common_infra.BuildASApiUri(1, "register")

	// Send the POST request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}

	return c.clientHandler.HandleResponse(resp)
}
