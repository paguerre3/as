package infrastructure

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

// integration test
func TestRegister(t *testing.T) {
	client := NewRegistrationClient()
	//response, statusCode, err := client.Register("aAguerre Pablo D.", "ARG", "aguerrepablodario@hotmail.com", "engineering")
	//response, statusCode, err := client.Register("PabloDariAguerre", "ARG", "p_d_aguerre@outlook.com", "engineering")
	response, statusCode, err := client.Register("Pablo_Dario_Aguerre", "ARG", "sejax32599@pokeline.com", "engineering")

	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response) // message says API key is sent via e-mail
}
