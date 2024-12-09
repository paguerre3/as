package application

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/paguerre3/as/internal/modules/1_stealth_probe/domain"
)

type ProbeUseCase interface {
	Execute() (response map[string]interface{}, statusCode int, err error)
}

type MeasurementClient interface {
	Measurement() (response map[string]interface{}, statusCode int, err error)
	MeasurementSolution(speed int) (response map[string]interface{}, statusCode int, err error)
} // Exposing duplicated interface to avoid DDD violations

type probeUseCaseImpl struct {
	measurementClient MeasurementClient
}

func NewProbeUseCase(measurementClient MeasurementClient) ProbeUseCase {
	return &probeUseCaseImpl{
		measurementClient: measurementClient,
	}
}

func (p *probeUseCaseImpl) Execute() (map[string]interface{}, int, error) {
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
				response, statusCode, err := p.measurementClient.Measurement()
				if err != nil || statusCode != 200 {
					errorChan <- fmt.Errorf("goroutine %d failed to fetch measurement: %v", id, err)
					return
				}

				log.Infof("Goroutine %d response: %+v", id, response)

				// Check if distance and time are both present
				distance, ok := response["distance"].(string)
				time, ok2 := response["time"].(string)

				if ok && ok2 && !strings.Contains(distance, "try again") {
					// If both distance and time are found, send them and cancel other goroutines
					resultChan <- struct {
						distance, time string
						err            error
					}{distance: distance, time: time, err: nil}
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

	// Wait for the first successful result or error, or timeout
	timeout := time.After(30 * time.Second) // Timeout to avoid deadlock
	select {
	case result := <-resultChan:
		// We have found both distance and time
		distance := result.distance
		t := result.time
		if distance == "" || t == "" {
			err := fmt.Errorf("distance or time is empty")
			log.Error(err)
			return nil, 0, err
		}

		// Proceed with speed calculation
		speed, err := domain.CalculateSpeed(distance, t)
		if err != nil {
			log.Errorf("error calculating speed: %v", err)
			return nil, 0, err
		}

		// Fetch solution after speed calculation = 405
		return p.measurementClient.MeasurementSolution(speed)
	case err := <-errorChan:
		// Handle any errors encountered in the goroutines
		log.Error(err)
		return nil, 0, err
	case <-timeout:
		// Handle timeout
		err := fmt.Errorf("operation timed out")
		log.Error(err)
		return nil, 0, err
	}
}
