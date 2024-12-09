package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	grid, enemyX, enemyY := ParseRadarData(LastRadarInfoAvailable)
	DisplayRadar(grid, enemyX, enemyY)
	assert.Equal(t, 8, len(grid))
	assert.Equal(t, "g", enemyX)
	assert.Equal(t, 5, enemyY)
}
