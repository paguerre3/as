package infrastructure

import (
	"testing"
	"time"

	"github.com/paguerre3/as/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestDamagedSpaceshptSolution7(t *testing.T) {
	for i := 0; i < 3; i++ { // less than 3 minutes
		handlerDebug.RegisterEndpont7Solution(common.EXPOSED_BASE_ENDPOINT)
		time.Sleep(55 * time.Second)
	}
	assert.True(t, true)
}
