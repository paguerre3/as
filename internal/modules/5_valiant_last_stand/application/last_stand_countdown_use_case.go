package application

import (
	"fmt"
	"time"

	"github.com/paguerre3/as/internal/modules/5_valiant_last_stand/domain"
)

type LastStandCountdownUseCase interface {
	Execute() (string, error)
}

type BattleActionsClient interface {
	StartBattle() (string, int, error)
	PerformTurn(action string, x string, y int) (map[string]interface{}, int, error)
} // Exposing duplicated interface to avoid DDD violations

type lastStandCountdownUseCaseImpl struct {
	battleActionsClient BattleActionsClient
}

func NewLastStandCountdownUseCase(battleActionsClient BattleActionsClient) LastStandCountdownUseCase {
	return &lastStandCountdownUseCaseImpl{
		battleActionsClient: battleActionsClient,
	}
}

func (l *lastStandCountdownUseCaseImpl) Execute() (string, error) {
	startMessage, statusCode, error := l.battleActionsClient.StartBattle()
	if error != nil {
		//return "", fmt.Errorf("error starting battle: %v", error)
	}
	if statusCode != 200 {
		//return "", fmt.Errorf("error starting battle: status code %d", statusCode)
	}
	fmt.Printf("Battle started: %s\n", startMessage)

	// last info available:
	radarData := domain.LastRadarInfoAvailable
	domain.ParseRadarData(radarData)
	grid, enemyX, enemyY := domain.ParseRadarData(radarData)
	domain.DisplayRadar(grid, enemyX, enemyY)

	for turnsRemaining := 4; turnsRemaining > 0; turnsRemaining-- {
		fmt.Printf("Turn %d, enemy at: %s%d\n", 5-turnsRemaining, enemyX, enemyY)

		// Leer radar (simulado con los datos iniciales)
		turnResult, statusCode, error := l.battleActionsClient.PerformTurn("radar", enemyX, enemyY)
		if error != nil {
			return "", fmt.Errorf("error performing turn: %v", error)
		}
		if statusCode != 200 {
			return "", fmt.Errorf("error performing turn: status code %d", statusCode)
		}
		performedAction, ok := turnResult["performed_action"]
		if !ok {
			return "", fmt.Errorf("error performing turn: missing performed_action")
		}
		/*turnsRemaining, ok := turnResult["turns_remaining"]
		if !ok {
			return "", fmt.Errorf("error performing turn: missing turns_remaining")
		}*/
		actionResult, ok := turnResult["action_result"]
		if !ok {
			return "", fmt.Errorf("error performing turn: missing action_result")
		}
		message, ok := turnResult["message"]
		if !ok {
			return "", fmt.Errorf("error performing turn: missing message")
		}

		turnResp := domain.TurnResponse{
			PerformedAction: performedAction.(string),
			ActionResult:    actionResult.(string),
			Message:         message.(string),
		}
		fmt.Printf("Turn result: %+v\n", turnResp)

		fmt.Println("Radar read (MESSAGE):", turnResp.Message)
		fmt.Println("Radar read (ACTION):", turnResp.ActionResult)

		if domain.IsRadarDataValid(turnResp.Message) {
			radarData = turnResp.Message
			grid, enemyX, enemyY = domain.ParseRadarData(radarData)
		} else if domain.IsRadarDataValid(turnResp.ActionResult) {
			radarData = turnResp.ActionResult
			grid, enemyX, enemyY = domain.ParseRadarData(radarData)
		}
		domain.DisplayRadar(grid, enemyX, enemyY)
		// avoid prediction here (this will be the enemy last position before hitting the friendly spaceship #)
		if enemyX == "b" && enemyY == 5 {
			// Atack in last movement
			attackResult, statusCode, error := l.battleActionsClient.PerformTurn("attack", "c", 7)
			if error != nil {
				return "", fmt.Errorf("error performing attack: %v", error)
			}
			if statusCode != 200 {
				return "", fmt.Errorf("error performing attack: status code %d", statusCode)
			}
			fmt.Printf("Attack result: %+v\n", attackResult)
			return "success", nil
		}

		// Predecir movimiento enemigo
		enemyX, enemyY = domain.SimpleEnemyPrediction(grid, enemyX, enemyY)
		time.Sleep(domain.RadarRefreshTime)
	}

	// Atack in last movement based on prediction
	attackResult, statusCode, error := l.battleActionsClient.PerformTurn("attack", enemyX, enemyY)
	if error != nil {
		return "", fmt.Errorf("error performing attack: %v", error)
	}
	if statusCode != 200 {
		return "", fmt.Errorf("error performing attack: status code %d", statusCode)
	}
	performedAction, ok := attackResult["performed_action"]
	if !ok {
		return "", fmt.Errorf("error performing attack: missing performed_action")
	}
	message, ok := attackResult["message"]
	if !ok {
		return "", fmt.Errorf("error performing attack: missing message")
	}
	fmt.Printf("Attack result: %s, Message: %s\n", performedAction.(string), message.(string))

	return message.(string), nil
}
