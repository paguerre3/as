package infrastructure

import (
	"fmt"
	"testing"

	"github.com/paguerre3/as/internal/domain"
	"github.com/stretchr/testify/assert"
)

// E5: La Última Defensa de la "Valiant" - ¡Cuenta Regresiva!
// Note: NO IA prediction (simplified), final result is done via debug seen in the console.
func TestLastDefenseSolution(t *testing.T) {
	result, error := domain.LastDefense(handler)
	assert.NoError(t, error)
	assert.NotEmpty(t, result)
	fmt.Printf("result: %s\n", result)
}
