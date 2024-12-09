package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateSpeed(t *testing.T) {
	tests := []struct {
		name           string
		distanceStr    string
		timeStr        string
		expectedSpeed  int
		expectedErrMsg string
	}{
		{
			name:           "valid input",
			distanceStr:    "535 AU",
			timeStr:        "1.3209876543209877 hours",
			expectedSpeed:  405,
			expectedErrMsg: "",
		},
		{
			name:           "valid input no suffixes",
			distanceStr:    "535",
			timeStr:        "1.3209876543209877",
			expectedSpeed:  405,
			expectedErrMsg: "",
		},
		{
			name:           "empty distance",
			distanceStr:    "",
			timeStr:        "2.5 hours",
			expectedSpeed:  0,
			expectedErrMsg: "distance or time is empty",
		},
		{
			name:           "empty time",
			distanceStr:    "10.5",
			timeStr:        "",
			expectedSpeed:  0,
			expectedErrMsg: "distance or time is empty",
		},
		{
			name:           "invalid distance",
			distanceStr:    "abc",
			timeStr:        "2.5",
			expectedSpeed:  0,
			expectedErrMsg: "error parsing distance",
		},
		{
			name:           "invalid time",
			distanceStr:    "10.5",
			timeStr:        "abc",
			expectedSpeed:  0,
			expectedErrMsg: "error parsing time",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			speed, err := CalculateSpeed(tt.distanceStr, tt.timeStr)
			if err != nil {
				assert.ErrorContains(t, err, tt.expectedErrMsg)
			}
			assert.Equal(t, tt.expectedSpeed, speed)
		})
	}
}
