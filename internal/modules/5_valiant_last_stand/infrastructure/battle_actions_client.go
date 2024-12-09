package infrastructure

import (
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

type BattleActionsClient interface {
	StartBattle() (string, int, error)
	PerformTurn(action string, x string, y int) (map[string]interface{}, int, error)
}

type battleActionsClientImpl struct {
	clientHandler common_infra.ClientHandler
}

func NewBattleActionsClient() BattleActionsClient {
	return &battleActionsClientImpl{
		clientHandler: common_infra.NewClientHandler(),
	}
}

func (c *battleActionsClientImpl) StartBattle() (string, int, error) {
	uri := common_infra.BuildASApiUri(1, "s1/e5/actions/start")
	resp, err := c.clientHandler.Client().R().
		SetHeader(common_infra.AUTHORIZATION, common_infra.API_KEY).
		Post(uri)
	if err != nil {
		return c.clientHandler.HandleStringError(resp, err)
	}
	return c.clientHandler.HandleStringResponse(resp)
}

func (c *battleActionsClientImpl) PerformTurn(action string, x string, y int) (map[string]interface{}, int, error) {
	requestBody := map[string]interface{}{
		"action": action,
		"attack_position": map[string]interface{}{
			"x": x,
			"y": y,
		},
	}
	uri := common_infra.BuildASApiUri(1, "s1/e5/actions/perform-turn")
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
