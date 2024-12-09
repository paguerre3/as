package infrastructure

import (
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type UserHiddenClient interface {
	UserAndPasswordSolution(username, password string) (map[string]interface{}, int, error)
}

type userHiddenImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewUserHiddenClient() UserHiddenClient {
	return &userHiddenImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *userHiddenImpl) UserAndPasswordSolution(username, password string) (map[string]interface{}, int, error) {
	requestBody := map[string]string{
		"username": username,
		"password": password,
	}
	uri := common_infra.BuildASApiUri(1, "s1/e4/solution")

	// Send the POST request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetHeader(common_infra.CONTENT_TYPE, common_infra.APPLICATION_JSON).
		SetBody(requestBody).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}
	return c.clientHandler.HandleResponse(resp)
}
