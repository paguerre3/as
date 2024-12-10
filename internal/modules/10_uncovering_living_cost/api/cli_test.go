package api

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
Use Go’s pprof for CPU and memory profiling:

go tool pprof galactic_cost_of_living cpu.prof
go tool pprof galactic_cost_of_living mem.prof
*/

// Enable profiling in code:
func init() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

// integration test
func RunCliTest(t *testing.T) {
	// for now, instead of enabling an entry point, we just run the CLI manually via testing
	err := RunCommandLine()
	assert.NoError(t, err)
}
