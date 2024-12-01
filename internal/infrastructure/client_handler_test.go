package infrastructure

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var (
	handler = NewClientHandler()
)

func verifyCorrectness(t *testing.T, response map[string]interface{}, statusCode int, err error) {
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response)

	resultValue := response["result"]
	assert.NotEmpty(t, resultValue)
	assert.Equal(t, "correct", resultValue)
}
