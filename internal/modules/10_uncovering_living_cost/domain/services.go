package domain

import (
	"github.com/uber/h3-go/v4"
)

const (
	resolution = 9 // between 0 (biggest cell) and 15 (smallest cell) -number suggested by the library
)

// Business logic for processing mobility data into H3-based features
type H3Service interface {
	GenerateH3Features(mobilityData []MobilityData) map[string]int
	CalculateH3Key(lat float64, lon float64) string
}
type h3ServiceImpl struct{}

func NewH3Service() H3Service {
	return &h3ServiceImpl{}
}

// GenerateH3Features processes mobility data into a map of H3 hexagon counts by "key"
func (h *h3ServiceImpl) GenerateH3Features(mobilityData []MobilityData) map[string]int {
	result := make(map[string]int)

	for _, record := range mobilityData {
		h3Key := h.CalculateH3Key(record.Lat, record.Lon)
		result[h3Key]++
	}

	return result
}

func (h *h3ServiceImpl) CalculateH3Key(lat float64, lon float64) string {
	h3Index := h3.LatLngToCell(h3.LatLng{Lat: lat, Lng: lon}, resolution)
	return h3Index.String()
}
