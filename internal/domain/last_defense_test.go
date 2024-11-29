package domain

import (
	"fmt"
	"testing"
)

func TestParseData(t *testing.T) {
	var radarData = "a01b01c01d01e01f01g01h01|a02b02c02d$2e02f02g02h02|a03b03c$3d03e03f03g03h03|a04b04c$4d04e04f04g04h04|a05b^5c05d05e05f05g05h05|a06b06c06d$6e06f06g06h06|a07b07c07d07e07f07g07h07|a08b08c08d08e#8f08g08h08|"
	parseRadarData(radarData)
	grid, enemyX, enemyY := parseRadarData(radarData)
	fmt.Println("Parsed Grid:")
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Printf("Enemy xy: %s%d", enemyX, enemyY)
	//assert.Equal(t, 8, len(grid))
	//assert.Equal(t, "g", enemyX)
	//assert.Equal(t, 5, enemyY)
}
