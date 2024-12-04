package domain

import (
	"fmt"
	"sync"

	"gonum.org/v1/gonum/stat"
)

var (
	repairCodes = map[string]string{
		"navigation":       "NAV-01",
		"communications":   "COM-02",
		"life_support":     "LIFE-03",
		"engines":          "ENG-04",
		"deflector_shield": "SHLD-05",
	}

	systems = []string{"navigation", "communications", "life_support", "engines", "deflector_shield"}
)

const (
	PICK_SYSTEM = "<pick one of the systems>"

	PressureCritical    = 10.0   // MPa
	TemperatureCritical = 500.0  // °C
	VolumeCritical      = 0.0035 // m^3/kg

	PressureLine = 0.05 // MPa (pressure where liquid and vapor lines begin)
)

type DamagedSpaceship interface {
	NextDamagedSystem() string
	RepairCode() (string, bool)
	SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error)
}

type damagedSpaceshipImpl struct {
	currentSystemIndex int
	mu                 sync.Mutex
	damagedSystem      string // current damaged system
}

func NewDamagedSpaceship() DamagedSpaceship {
	return &damagedSpaceshipImpl{
		currentSystemIndex: 0,
		mu:                 sync.Mutex{},
		damagedSystem:      PICK_SYSTEM,
	}
}

func (d *damagedSpaceshipImpl) NextDamagedSystem() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.damagedSystem = systems[d.currentSystemIndex]
	d.currentSystemIndex = (d.currentSystemIndex + 1) % len(systems)
	return d.damagedSystem
}

func (d *damagedSpaceshipImpl) RepairCode() (string, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	code, ok := repairCodes[d.damagedSystem]
	return code, ok
}

// Saturated liquid line equation
func (d *damagedSpaceshipImpl) SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error) {
	// Check input range
	if pressure < PressureLine || pressure > PressureCritical {
		return 0, 0, fmt.Errorf("pressure %.2f MPa is out of range (%.2f MPa - %.2f MPa)", pressure, PressureLine, PressureCritical)
	}

	// Example data points (replace with empirical values or formulas)
	pressures := []float64{PressureLine, PressureCritical}
	volumesLiquid := []float64{0.00105, VolumeCritical}
	volumesVapor := []float64{30.0, VolumeCritical}

	// Linear regression for specific volumes
	liquidAlpha, liquidBeta := stat.LinearRegression(pressures, volumesLiquid, nil, false)
	vaporAlpha, vaporBeta := stat.LinearRegression(pressures, volumesVapor, nil, false)

	// Predict specific volumes using the regression equations
	liquidVolume := liquidAlpha + liquidBeta*pressure
	vaporVolume := vaporAlpha + vaporBeta*pressure

	return liquidVolume, vaporVolume, nil
}
