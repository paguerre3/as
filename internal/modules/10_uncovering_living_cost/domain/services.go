package domain

import (
	"sync"

	"github.com/uber/h3-go/v4"
)

// Business logic for processing mobility data into H3-based features
type H3Service interface {
	GenerateH3Features(mobilityData []MobilityData, resolution int) map[string]int
}
type h3ServiceImpl struct{}

func NewH3Service() H3Service {
	return &h3ServiceImpl{}
}

// GenerateH3Features processes mobility data into a map of H3 hexagon counts
func (h *h3ServiceImpl) GenerateH3Features(mobilityData []MobilityData, resolution int) map[string]int {
	h3Map := sync.Map{}
	var wg sync.WaitGroup

	chunkSize := len(mobilityData) / 4
	for i := 0; i < len(mobilityData); i += chunkSize {
		end := i + chunkSize
		if end > len(mobilityData) {
			end = len(mobilityData)
		}

		wg.Add(1)
		go func(chunk []MobilityData) {
			defer wg.Done()
			localMap := make(map[string]int)
			for _, record := range chunk {
				// 1st strategy uses only latitude and longitude:
				h3Index := h3.LatLngToCell(h3.LatLng{Lat: record.Lat, Lng: record.Lon}, resolution)
				localMap[h3Index.String()]++
			}
			for k, v := range localMap {
				h3Map.LoadOrStore(k, v)
				if count, ok := h3Map.Load(k); ok {
					h3Map.Store(k, count.(int)+v)
				}
			}
		}(mobilityData[i:end])
	}

	wg.Wait()

	// Convert sync.Map to regular map for use
	result := make(map[string]int)
	h3Map.Range(func(key, value interface{}) bool {
		result[key.(string)] = value.(int)
		return true
	})
	return result
}
