package infrastructure

import (
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/domain"
	"github.com/stretchr/testify/assert"
)

// E6: La Infiltración en Ciudad Prisma: Un Desafío para los Maestros de Datos (Pokemons)
func TestPokemonsHeightAvgSolution(t *testing.T) {
	solution, error := domain.CalculatePokemonTypesAverageHeights(handler)
	assert.NoError(t, error)
	assert.NotEmpty(t, solution)
	log.Infof("solution: %s\n", solution)
	response, statusCode, err := handler.PokemonSolution(solution)
	verifyCorrectness(t, response, statusCode, err)
}
