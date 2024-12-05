package domain

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
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

	// Critical values
	CriticalPressureMPa  = 10.0   // Critical Pressure in MPa
	CriticalTemperatureC = 500.0  // Critical Temperature in °C
	CriticalVolume       = 0.0035 // Critical Volume in m^3/kg

	// Reference and starting values
	ReferencePressureMPa = 0.05    // Reference Pressure in MPa
	LiquidStartVolume    = 0.00105 // Saturated Liquid Volume at 0.05 MPa
	VaporStartVolume     = 30.00   // Saturated Vapor Volume at 0.05 MPa
	PressureLine         = 0.01    // Minimum pressure line (adjustable based on needs)
	PressureCritical     = 10.0    // Critical pressure
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

// SaturatedLiquidAndVaporVolumes calculates the saturated liquid and vapor volumes for a given pressure using interpolation
func (d *damagedSpaceshipImpl) SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error) {
	// Check input range
	if pressure < PressureLine || pressure > PressureCritical {
		return 0, 0, fmt.Errorf("pressure %.2f MPa is out of range (%.2f MPa - %.2f MPa)", pressure, PressureLine, PressureCritical)
	}

	// Empirical data points for liquid and vapor volumes at different pressures
	pressures := []float64{ReferencePressureMPa, CriticalPressureMPa}
	volumesLiquid := []float64{LiquidStartVolume, CriticalVolume}
	volumesVapor := []float64{VaporStartVolume, CriticalVolume}

	// Linear interpolation for specific volumes
	liquidVolume := interpolate(pressures, volumesLiquid, pressure)
	vaporVolume := interpolate(pressures, volumesVapor, pressure)

	liquidVolumeRounded, _ := decimal.NewFromFloat(liquidVolume).Round(5).Float64()
	vaporVolumeRounded, _ := decimal.NewFromFloat(vaporVolume).Round(5).Float64()

	return liquidVolumeRounded, vaporVolumeRounded, nil
}

// interpolate performs linear interpolation for a given x value
func interpolate(xs, ys []float64, x float64) float64 {
	for i := 1; i < len(xs); i++ {
		if x <= xs[i] {
			x0, x1 := xs[i-1], xs[i]
			y0, y1 := ys[i-1], ys[i]
			return y0 + (y1-y0)*(x-x0)/(x1-x0)
		}
	}
	return ys[len(ys)-1]
}
