package domain

import (
	"fmt"
	"strings"
	"time"
)

const (
	gridSize         = 8
	radarRefreshTime = 1 * time.Second // Simulación de tiempo para refrescar
)

type TurnResponse struct {
	PerformedAction string `json:"performed_action"`
	TurnsRemaining  int    `json:"turns_remaining"`
	TimeRemaining   int    `json:"time_remaining"`
	ActionResult    string `json:"action_result"`
	Message         string `json:"message"`
}

func parseRadarData(radarData string) ([][]string, string, int) {
	// Define the grid size for an 8x8 space
	grid := make([][]string, gridSize)

	// Split the radar data into rows based on '|'
	rows := strings.Split(radarData, "|")

	var enemyX string
	var enemyY int

	// Loop through rows and parse columns
	for y, row := range rows {
		if row == "" {
			continue
		}
		grid[y] = make([]string, gridSize)
		for x := 0; x < len(row); x += 3 {
			// Extract the 3-character cell
			cell := row[x : x+3]
			grid[y][x/3] = cell

			// Check if the cell contains the enemy character '^'
			if strings.Contains(cell, "^") {
				enemyX = string(cell[0])    // Extract the column (letter)
				enemyY = int(cell[2] - '0') // Extract the row (number)
			}
		}
	}

	return grid, enemyX, enemyY
}

func predictEnemyMovement(grid [][]string, enemyX string, enemyY int) (string, int) {
	columnMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
	reverseColumnMap := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h"}

	// Convert enemyX to a column index
	enemyColumnIdx := columnMap[enemyX]

	// Find the target position (#)
	targetX, targetY := -1, -1
	for y := 0; y < gridSize; y++ {
		for x := 1; x <= gridSize; x++ { // Use 1-based column indices
			if strings.Contains(grid[y][x-1], "#") { // Adjust index for zero-based slice access
				targetX, targetY = x, y
				break
			}
		}
		if targetX != -1 {
			break
		}
	}

	// If no target (#) is found, return the current position
	if targetX == -1 || targetY == -1 {
		return enemyX, enemyY
	}

	// Predict movement towards the target while avoiding obstacles ($)
	if enemyY < targetY && grid[enemyY][enemyColumnIdx-1] != "$" {
		enemyY++ // Move down
	} else if enemyColumnIdx < targetX && grid[enemyY][enemyColumnIdx] != "$" {
		enemyColumnIdx++ // Move right
	} else if enemyColumnIdx > targetX && grid[enemyY][enemyColumnIdx-2] != "$" {
		enemyColumnIdx-- // Move left
	}

	// Convert the column index back to a letter
	enemyX = reverseColumnMap[enemyColumnIdx]
	fmt.Printf("PREDICTED Enemy Position XY: %s%d\n", enemyX, enemyY)
	return enemyX, enemyY
}

func displayRadar(grid [][]string, enemyX string, enemyY int) {
	fmt.Println("Parsed Grid:")
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Printf("Actual Enemy Position XY: %s%d\n", enemyX, enemyY)
}

func isRadarDataValid(message string) bool {
	return strings.Contains(message, "^") && strings.Contains(message, "|")
}

func LastDefense(handler ClientHandler) (string, error) {
	startMessage, statusCode, error := handler.StartBattle()
	if error != nil {
		//return "", fmt.Errorf("error starting battle: %v", error)
	}
	if statusCode != 200 {
		//return "", fmt.Errorf("error starting battle: status code %d", statusCode)
	}
	fmt.Printf("Battle started: %s\n", startMessage)

	// last info available:
	radarData := "a01b01c01d01e01f01g01h01|a02b02c02d02e$2f02g02h02|a03b03c03d03e03f03g03h$3|a04b04c04d04e04f04g04h04|a05b05c05d05e$5f05g^5h05|a06b06c06d06e$6f06g06h06|a07b07c07d07e07f07g07h07|a08b08c08d08e08f#8g08h08|"
	parseRadarData(radarData)
	grid, enemyX, enemyY := parseRadarData(radarData)
	displayRadar(grid, enemyX, enemyY)

	for turnsRemaining := 4; turnsRemaining > 0; turnsRemaining-- {
		fmt.Printf("Turn %d, enemy at: %s%d\n", 5-turnsRemaining, enemyX, enemyY)

		// Leer radar (simulado con los datos iniciales)
		turnResult, statusCode, error := handler.PerformTurn("radar", enemyX, enemyY)
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

		turnResp := TurnResponse{
			PerformedAction: performedAction.(string),
			ActionResult:    actionResult.(string),
			Message:         message.(string),
		}
		fmt.Printf("Turn result: %+v\n", turnResp)

		fmt.Println("Radar read (MESSAGE):", turnResp.Message)
		fmt.Println("Radar read (ACTION):", turnResp.ActionResult)

		if isRadarDataValid(turnResp.Message) {
			radarData = turnResp.Message
			grid, enemyX, enemyY = parseRadarData(radarData)
		} else if isRadarDataValid(turnResp.ActionResult) {
			radarData = turnResp.ActionResult
			grid, enemyX, enemyY = parseRadarData(radarData)
		}
		displayRadar(grid, enemyX, enemyY)
		// avoid prediction here
		if enemyX == "b" || enemyY == 5 {
			// Atack in last movement
			attackResult, statusCode, error := handler.PerformTurn("attack", "c", 7)
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
		enemyX, enemyY = predictEnemyMovement(grid, enemyX, enemyY)
		time.Sleep(radarRefreshTime)
	}

	// Atack in last movement based on prediction
	attackResult, statusCode, error := handler.PerformTurn("attack", enemyX, enemyY)
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
