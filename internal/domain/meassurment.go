package domain

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// CalculateSpeed calculates the speed from distance and time strings.
// It expects distance in AU (astronomical units) and time in hours (both with or without suffixes).
func CalculateSpeed(distanceStr, timeStr string) (int, error) {
	if distanceStr == "" || timeStr == "" {
		return 0, fmt.Errorf("distance or time is empty")
	}
	distanceStr = cleanString(distanceStr)
	timeStr = cleanString(timeStr)

	// Convert the cleaned distance and time strings to float64
	distance, err := strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing distance: %v", err)
	}

	time, err := strconv.ParseFloat(timeStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing time: %v", err)
	}

	// Calculate speed (in AU per hour)
	speed := distance / time

	// Round the speed to the nearest integer
	roundedSpeed := math.Round(speed)

	// Return the speed as an integer
	return int(roundedSpeed), nil
}

// Helper function to clean a string by removing suffixes or units (e.g., "AU", "hours")
func cleanString(input string) string {
	// Remove spaces and handle suffixes if they exist
	input = strings.TrimSpace(input)
	// Remove known suffixes if present
	suffixes := []string{" AU", " hours"}
	for _, suffix := range suffixes {
		if strings.HasSuffix(input, suffix) {
			input = strings.TrimSuffix(input, suffix)
			break
		}
	}
	return input
}
