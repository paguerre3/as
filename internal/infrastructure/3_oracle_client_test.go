package infrastructure

import (
	"strings"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/domain"
	"github.com/stretchr/testify/assert"
)

// E3: La Búsqueda del Templo Sith Perdido
func TestBalancedPlanetSolution(t *testing.T) {
	// Actual test logic goes here:
	planets, err := domain.AllPlanets(handler)
	assert.NoError(t, err)
	assert.NotEmpty(t, planets)

	for _, planet := range planets {
		ibf, err := domain.CalculateIBF(handler, planet)
		if err != nil {
			// only possible error is "no residents found" whihc produces 0 IBF
			log.Warnf("Error calculating IBF for planet %s: %v", planet.Name, err)
		}
		if ibf == 0 && err == nil {
			// only one panet with people and balanced (IBF = 0)
			response, statusCode, err := handler.OracleSolution(planet.Name)
			verifyCorrectness(t, response, statusCode, err)
			log.Infof(strings.Repeat("-", 75))
			log.Infof(strings.Repeat("-", 75))
			// "Balanced Planet: Ryloth"
			log.Infof("Balanced Planet: %s", planet.Name)
			log.Infof(strings.Repeat("-", 75))
		}
	}
}
