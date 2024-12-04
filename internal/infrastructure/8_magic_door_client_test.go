package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMagicDoors(t *testing.T) {
	response, statusCode, err := handlerDebug.FirstClues()
	assert.NoError(t, err)
	assert.True(t, len(response) > 0)
	assert.Equal(t, 200, statusCode)
	response, statusCode, err = handlerDebug.HiddenMessageSolution(response)
	verifyCorrectness(t, response, statusCode, err)
}
