package domain

import (
	"fmt"
	"math"
	"strconv"
)

func CalculateSpeed(distanceStr, timeStr string) (int, error) {
	if distanceStr == "" || timeStr == "" {
		return 0, fmt.Errorf("distance or time is empty")
	}

	// Step 1: Convert the distance and time to float64
	distance, err := strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing distance: %v", err)
	}

	time, err := strconv.ParseFloat(timeStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing time: %v", err)
	}

	// Step 2: Calculate speed (in AU per hour)
	speed := distance / time

	// Step 3: Round the speed to the nearest integer
	roundedSpeed := math.Round(speed)
	return int(roundedSpeed), nil
}
