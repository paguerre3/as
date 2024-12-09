package domain

func CalculateAverageHeights(typeHeights map[string][]float64) map[string]float64 {
	averageHeights := make(map[string]float64)
	for typeName, heights := range typeHeights {
		var totalHeights float64
		for _, height := range heights {
			totalHeights += height
		}
		averageHeights[typeName] = float64(totalHeights) / float64(len(heights))
	}
	return averageHeights
}
