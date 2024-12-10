package infrastructure

import (
	"testing"

	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

// E4: La Búsqueda de la Forja Élfica Olvidada
// integration test
func TestUserAndPasswordSolution(t *testing.T) {
	//<span style="color:white; float: right;">Not all those who wander</span>
	//<input type="password" value="are lost" readonly="">
	userHiddenClient := NewUserHiddenClient()
	response, statusCode, err := userHiddenClient.UserAndPasswordSolution("Not all those who wander", "are lost")
	common_infra.VerifyCorrectness(t, response, statusCode, err)
}
