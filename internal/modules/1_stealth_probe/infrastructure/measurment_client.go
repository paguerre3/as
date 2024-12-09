package infrastructure

import (
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type MeasurementClient interface {
	Measurement() (response map[string]interface{}, statusCode int, err error)
	MeasurementSolution(speed int) (response map[string]interface{}, statusCode int, err error)
}

type measurementClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewMeasurementClient() MeasurementClient {
	return &measurementClientImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *measurementClientImpl) Measurement() (map[string]interface{}, int, error) {
	uri := common_infra.BuildASApiUri(1, "s1/e1/resources/measurement")

	// Send the GET request
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		Get(uri)
	if err != nil {
		return c.clientHandler.HandleError(resp, err)
	}

	return c.clientHandler.HandleResponse(resp)
}

func (c *measurementClientImpl) MeasurementSolution(speed int) (response map[string]interface{}, statusCode int, err error) {
	requestBody := map[string]int{
		"speed": speed,
	}

	uri := common_infra.BuildASApiUri(1, "s1/e1/solution")

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
