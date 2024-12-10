package infrastructure

import (
	"testing"

	"github.com/paguerre3/as/internal/modules/8_magic_door_spell/application"
	"github.com/paguerre3/as/internal/modules/8_magic_door_spell/domain"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

// integration test
func TestMagicDoors(t *testing.T) {
	magicDoorClient := NewMagicDoorClient(domain.NewGryffindorCookie())
	magicDoorUseCase := application.NewMagicDoorUseCase(magicDoorClient)
	response, statusCode, err := magicDoorUseCase.Execute()
	common_infra.VerifyCorrectness(t, response, statusCode, err)
}
