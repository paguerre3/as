package domain

import (
	"encoding/json"
	"fmt"
)

type Star struct {
	ID        string  `json:"id"`
	Resonance float64 `json:"resonance"`
	Position  struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"position"`
}

func CalculateAverageResonance(handler ClientHandler) (int, error) {
	var totalResonance, count float64
	page := 1

	for {
		response, statusCode, err := handler.FetchStars(page)
		if err != nil {
			return 0, err
		}
		if statusCode != 200 {
			return 0, fmt.Errorf("status code: %d", statusCode)
		}
		stars, err := convertToStars(response)
		if err != nil {
			return 0, err
		}
		if len(stars) == 0 {
			break
		}

		for _, star := range stars {
			totalResonance += star.Resonance
			count++
		}
		page++
	}

	return int(totalResonance / count), nil
}

func convertToStars(data []map[string]interface{}) (stars []Star, err error) {
	if len(data) == 0 {
		// data is empty
		return nil, nil
	}
	for _, item := range data {
		// Marshal the map to JSON bytes
		jsonData, err := json.Marshal(item)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal map: %v", err)
		}

		// Unmarshal the JSON bytes into the Star struct
		var star Star
		if err := json.Unmarshal(jsonData, &star); err != nil {
			return nil, fmt.Errorf("failed to unmarshal into Star: %v", err)
		}

		stars = append(stars, star)
	}
	return stars, nil
}
