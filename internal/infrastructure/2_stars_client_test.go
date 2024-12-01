package infrastructure

import (
	"testing"

	"github.com/paguerre3/as/internal/domain"
	"github.com/stretchr/testify/assert"
)

// E2: El Enigma Cósmico de Kepler-452b
func TestFetchStarsAndResonanceSolution(t *testing.T) {
	avg, err := domain.CalculateAverageResonance(handler)
	assert.NoError(t, err)
	assert.NotZero(t, avg)

	response, statusCode, err := handler.ResonanceSolution(avg) // 388
	verifyCorrectness(t, response, statusCode, err)
}
