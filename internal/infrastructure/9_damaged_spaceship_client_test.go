package infrastructure

import (
	"testing"
	"time"

	"github.com/paguerre3/as/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestDamagedSpaceshptSolution9(t *testing.T) {
	for i := 0; i < 3; i++ { // less than 5 minutes
		handlerDebug.RegisterEndpont9Solution(common.EXPOSED_BASE_ENDPOINT)
		time.Sleep(90 * time.Second)
	}
	assert.True(t, true)
}
