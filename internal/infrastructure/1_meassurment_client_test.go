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
		verifyCorrectness(t, response, statusCode, err)
	case err := <-errorChan:
		// Handle any errors encountered in the goroutines
		t.Fatal(err)
	}
}
