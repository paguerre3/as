package infrastructure

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	//response, statusCode, err := handler.Register("DarioAguerre", "ARG", "aguerrepablodario@gmail.com", "engineering")
	//response, statusCode, err := handler.Register("aAguerre Pablo D.", "ARG", "aguerrepablodario@hotmail.com", "engineering")
	response, statusCode, err := handler.Register("PabloDariAguerre", "ARG", "p_d_aguerre@outlook.com", "engineering")

	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response) // message says API key is sent via e-mail

	/** e-mail received:
	Esta es tu API-KEY para la AltScore Contest
	//255292ff68394c6eb6136069a034bf28
	//36b39441ba7040fea158db6a4103aaa6
	20adafe2bf3a4a43b20792aceac3dad9
	*/
}
