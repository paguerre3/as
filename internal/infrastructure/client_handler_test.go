package infrastructure

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	handler = NewClientHandler()
)

func TestRegister(t *testing.T) {
	response, statusCode, err := handler.Register("DarioAguerre", "ARG", "aguerrepablodario@gmail.com", "engineering")
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response) // message says API key is sent via e-mail

	/** e-mail received:
	Esta es tu API-KEY para la AltScore Contest
	255292ff68394c6eb6136069a034bf28
	*/
}

// E1: La Sonda Silenciosa
func TestMeasurementAndSolution(t *testing.T) {
	// Create a cancelable context to handle goroutine cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channels for receiving results or errors
	resultChan := make(chan struct {
		distance, time string
		err            error
	}, 1)
	errorChan := make(chan error, 1)

	// Number of goroutines to run concurrently (goroutine pool size)
	numGoroutines := 50

	// Function to fetch distance and time in each goroutine
	fetchMeasurement := func(ctx context.Context, id int) {
		for {
			select {
			case <-ctx.Done():
				// If context is canceled, exit the goroutine
				return
			default:
				response, statusCode, err := handler.Measurement()
				if err != nil || statusCode != 200 {
					errorChan <- fmt.Errorf("goroutine %d failed to fetch measurement: %v", id, err)
					return
				}

				log.Infof("Goroutine %d response: %+v", id, response)

				// Check if distance and time are both present
				distance, ok := response["distance"]
				time, ok2 := response["time"]

				if (ok && ok2) && !strings.Contains(distance.(string), "try again") {
					// If both distance and time are found, send them and cancel other goroutines
					resultChan <- struct {
						distance, time string
						err            error
					}{distance: distance.(string), time: time.(string), err: nil}
					cancel() // Cancel the context to stop other goroutines
					return
				}
			}

			// Delay to prevent too frequent retries (optional)
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Launch the goroutines
	for i := 0; i < numGoroutines; i++ {
		go fetchMeasurement(ctx, i)
	}

	// Wait for the first successful result or error
	select {
	case result := <-resultChan:
		// We have found both distance and time
		distance := result.distance
		time := result.time
		assert.NotEmpty(t, distance)
		assert.NotEmpty(t, time)

		// Proceed with speed calculation
		speed, err := domain.CalculateSpeed(distance, time)
		assert.NoError(t, err)

		// Fetch solution after speed calculation
		response, statusCode, err := handler.MeassurmentSolution(speed) // 405
		assert.NoError(t, err)
		assert.Equal(t, 200, statusCode)
		log.Infof("response: %+v", response)

		resultValue := response["result"]
		assert.NotEmpty(t, resultValue)
		assert.Equal(t, "correct", resultValue)
	case err := <-errorChan:
		// Handle any errors encountered in the goroutines
		t.Fatal(err)
	}
}

// E2: El Enigma Cósmico de Kepler-452b
func TestFetchStarsAndResonanceSolution(t *testing.T) {
	avg, err := domain.CalculateAverageResonance(handler)
	assert.NoError(t, err)
	assert.NotZero(t, avg)

	response, statusCode, err := handler.ResonanceSolution(avg) // 388
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response)

	resultValue := response["result"]
	assert.NotEmpty(t, resultValue)
	assert.Equal(t, "correct", resultValue)
}

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
			assert.NoError(t, err)
			assert.Equal(t, 200, statusCode)
			log.Infof("response: %+v", response)

			resultValue := response["result"]
			assert.NotEmpty(t, resultValue)
			assert.Equal(t, "correct", resultValue)
			log.Infof(strings.Repeat("-", 75))
			log.Infof(strings.Repeat("-", 75))
			// "Balanced Planet: Ryloth"
			log.Infof("Balanced Planet: %s", planet.Name)
			log.Infof(strings.Repeat("-", 75))
		}
	}
}

// E4: La Búsqueda de la Forja Élfica Olvidada
func TestUserAndPasswordSolution(t *testing.T) {
	//<span style="color:white; float: right;">Not all those who wander</span>
	//<input type="password" value="are lost" readonly="">
	response, statusCode, err := handler.UserAndPasswordSolution("Not all those who wander", "are lost")
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response)

	resultValue := response["result"]
	assert.NotEmpty(t, resultValue)
	assert.Equal(t, "correct", resultValue)
}

// E5: La Última Defensa de la "Valiant" - ¡Cuenta Regresiva!
func TestLastDefenseSolution(t *testing.T) {
	result, error := domain.LastDefense(handler)
	assert.NoError(t, error)
	assert.NotEmpty(t, result)
	fmt.Printf("result: %s\n", result)
}

// E6: La Infiltración en Ciudad Prisma: Un Desafío para los Maestros de Datos (Pokemons)
func TestPokemonsHeightAvgSolution(t *testing.T) {
	solution, error := domain.CalculatePokemonTypesAverageHeights(handler)
	assert.NoError(t, error)
	assert.NotEmpty(t, solution)
	fmt.Printf("solution: %s\n", solution)
	response, statusCode, err := handler.PokemonSolution(solution)
	assert.NoError(t, err)
	assert.Equal(t, 200, statusCode)
	log.Infof("response: %+v", response)
	resultValue := response["result"]
	assert.NotEmpty(t, resultValue)
	assert.Equal(t, "correct", resultValue)
}