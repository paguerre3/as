package application

import "github.com/paguerre3/as/internal/modules/10_uncovering_living_cost/domain"

type H3UseCase interface {
	GenerateFeatures(mobilityData []domain.MobilityData) map[string]int
	CalculateH3Key(lat float64, lon float64) string
}

type h3UseCaseImpl struct {
	h3Service domain.H3Service
}

// Wrapper for domain H3Service to prepare for DDD structure
func NewH3UseCase() H3UseCase {
	return &h3UseCaseImpl{
		h3Service: domain.NewH3Service(),
	}
}

// GenerateFeatures delegates to the domain H3Service
func (h *h3UseCaseImpl) GenerateFeatures(mobilityData []domain.MobilityData) map[string]int {
	return h.h3Service.GenerateH3Features(mobilityData)
}

func (h *h3UseCaseImpl) CalculateH3Key(lat float64, lon float64) string {
	return h.h3Service.CalculateH3Key(lat, lon)
}
