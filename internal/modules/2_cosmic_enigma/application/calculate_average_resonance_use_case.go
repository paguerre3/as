package application

import (
	"fmt"

	"github.com/paguerre3/as/internal/modules/2_cosmic_enigma/domain"
)

type CalculateAverageResonanceUseCase interface {
	Execute() (response map[string]interface{}, statusCode int, err error)
}

type ResonanceClient interface {
	FetchStars(page int) (response []map[string]interface{}, statusCode int, err error)
	ResonanceSolution(averaegeResonance int) (response map[string]interface{}, statusCode int, err error)
} // Exposing duplicated interface to avoid DDD violations

type calculateAverageResonanceUseCaseImpl struct {
	resonanceClient ResonanceClient
}

func NewCalculateAverageResonanceUseCase(resonanceClient ResonanceClient) CalculateAverageResonanceUseCase {
	return &calculateAverageResonanceUseCaseImpl{
		resonanceClient: resonanceClient,
	}
}

func (c *calculateAverageResonanceUseCaseImpl) Execute() (map[string]interface{}, int, error) {
	var totalStarsResonance, starsCount float64
	page := 1
	for {
		response, statusCode, err := c.resonanceClient.FetchStars(page)
		if err != nil {
			return nil, 0, err
		}
		if statusCode != 200 {
			return nil, 0, fmt.Errorf("status code: %d", statusCode)
		}
		stars, err := domain.ConvertToStars(response)
		if err != nil {
			return nil, 0, err
		}
		if len(stars) == 0 {
			break
		}

		for _, star := range stars {
			totalStarsResonance += star.Resonance
			starsCount++
		}
		page++
	}
	// 388
	avg := domain.AverageResonance(totalStarsResonance, starsCount)
	return c.resonanceClient.ResonanceSolution(avg)
}
