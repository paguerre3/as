package domain

import "fmt"

type Planet struct {
	Name      string   `json:"name"`
	Residents []string `json:"residents"`
}

// Índice de Balance de la Fuerza:
/*
IBF = ((Número de Personajes del Lado Luminoso) - (Número de Personajes del Lado Oscuro)) /
       (Total de Personajes en el Planeta)

When IBF is 0 and "there are residents" then the planet is balanced.
*/
func CalculateIBF(lightSideCount, darkSideCount int) (float64, error) {
	// Calcular el IBF
	totalCount := lightSideCount + darkSideCount
	if totalCount == 0 {
		return 0, fmt.Errorf("no residents found")
	}
	ibf := float64(lightSideCount-darkSideCount) / float64(totalCount)
	return ibf, nil
}
