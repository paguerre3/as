package infrastructure

import (
	"testing"
)

// E4: La Búsqueda de la Forja Élfica Olvidada
func TestUserAndPasswordSolution(t *testing.T) {
	//<span style="color:white; float: right;">Not all those who wander</span>
	//<input type="password" value="are lost" readonly="">
	response, statusCode, err := handler.UserAndPasswordSolution("Not all those who wander", "are lost")
	verifyCorrectness(t, response, statusCode, err)
}
