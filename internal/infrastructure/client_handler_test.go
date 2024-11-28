package infrastructure

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	handler = NewClientHandler()
)

func TestRegister(t *testing.T) {
	response, statusCode, err := handler.Register("CamiAguerre", "ARG", "pablo.aguerre@gmail.com", "engineering")
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response)

	/**
	Esta es tu API-KEY para la AltScore Contest
	a79f99a48ee04b529605b797fe43182c
	*/
}

func TestMeasurementAndSolution(t *testing.T) {
	distance := ""
	time := ""
	for {
		response, statusCode, err := handler.Measurement()
		assert.NoError(t, err)
		assert.Equal(t, 200, statusCode)
		log.Infof("response: %+v", response)
		distance, ok := response["distance"]
		if !ok {
			continue
		}
		assert.NotEmpty(t, distance)
		time, ok := response["time"]
		if !ok {
			continue
		}
		assert.NotEmpty(t, time)
		break
	}
	speed, err := domain.CalculateSpeed(distance, time)
	assert.NoError(t, err)
	response, statusCode, err := handler.MeassurmentSolution(speed)
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response)
	result := response["result"]
	assert.NotEmpty(t, result)
}
