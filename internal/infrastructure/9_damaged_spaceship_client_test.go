package infrastructure

import (
	"testing"
	"time"

	"github.com/paguerre3/as/internal/common"
)

func TestDamagedSpaceshptSolution9(t *testing.T) {
	for i := 0; i < 3; i++ { // less than 5 minutes
		handlerDebug.RegisterEndpont9Solution(common.EXPOSED_BASE_ENDPOINT)
		// Pausar por 1.5 segundos (1500 milisegundos)
		time.Sleep(1500 * time.Millisecond)
	}
}
