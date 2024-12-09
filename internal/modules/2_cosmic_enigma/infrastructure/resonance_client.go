package infrastructure

import (
	"fmt"

	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type ResonanceClient interface {
	FetchStars(page int) (response []map[string]interface{}, statusCode int, err error)
	ResonanceSolution(averaegeResonance int) (response map[string]interface{}, statusCode int, err error)
}

type resonanceClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewResonanceClient() ResonanceClient {
	return &resonanceClientImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *resonanceClientImpl) FetchStars(page int) ([]map[string]interface{}, int, error) {
	uri := common_infra.BuildASApiUri(1, "s1/e2/resources/stars")

	// Send the GET request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		SetQueryParam("page", fmt.Sprintf("%d", page)).
		Get(uri)
	if err != nil {
		return c.clientHandler.HandleArrayError(resp, err)
	}

	return c.clientHandler.HandleArrayResponse(resp)
}

func (c *resonanceClientImpl) ResonanceSolution(averageResonance int) (map[string]interface{}, int, error) {
	requestBody := map[string]int{
		"average_resonance": averageResonance,
	}

	uri := common_infra.BuildASApiUri(1, "s1/e2/solution")

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
