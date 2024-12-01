package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextDamagedSystemAndRepairCode(t *testing.T) {
	ds := NewDamagedSpaceship()
	repairCode, ok := ds.RepairCode()
	assert.False(t, ok)
	assert.Equal(t, "", repairCode)

	tests := []struct {
		name               string
		expectedSystem     string
		expectedRepairCode string
	}{
		{
			name:               "first system",
			expectedSystem:     "navigation",
			expectedRepairCode: "NAV-01",
		},
		{
			name:               "second system",
			expectedSystem:     "communications",
			expectedRepairCode: "COM-02",
		},
		{
			name:               "third system",
			expectedSystem:     "life_support",
			expectedRepairCode: "LIFE-03",
		},
		{
			name:               "fourth system",
			expectedSystem:     "engines",
			expectedRepairCode: "ENG-04",
		},
		{
			name:               "fifth system",
			expectedSystem:     "deflector_shield",
			expectedRepairCode: "SHLD-05",
		},
		{
			name:               "wrap around to first system",
			expectedSystem:     "navigation", // start again from begining
			expectedRepairCode: "NAV-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSystem := ds.NextDamagedSystem()
			assert.Equal(t, tt.expectedSystem, actualSystem)
			repairCode, ok := ds.RepairCode()
			assert.True(t, ok)
			assert.Equal(t, tt.expectedRepairCode, repairCode)
		})
	}
}
