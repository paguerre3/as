package domain

import (
	"fmt"
	"sync"

	"gonum.org/v1/gonum/interp"
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

func (d *damagedSpaceshipImpl) SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error) {
	if pressure < PressureLine || pressure > PressureCritical {
		return 0, 0, fmt.Errorf("pressure %.2f MPa is out of range (%.2f MPa - %.2f MPa)", pressure, PressureLine, PressureCritical)
	}

	pressures := []float64{ReferencePressureMPa, CriticalPressureMPa}
	liquidVolumes := []float64{LiquidStartVolume, CriticalVolume}
	vaporVolumes := []float64{VaporStartVolume, CriticalVolume}

	liquidSpline := interp.AkimaSpline{}
	err := liquidSpline.Fit(pressures, liquidVolumes)
	if err != nil {
		return 0, 0, fmt.Errorf("error fitting liquid spline: %v", err)
	}

	vaporSpline := interp.AkimaSpline{}
	err = vaporSpline.Fit(pressures, vaporVolumes)
	if err != nil {
		return 0, 0, fmt.Errorf("error fitting vapor spline: %v", err)
	}

	liquidVolume := liquidSpline.Predict(pressure)
	vaporVolume := vaporSpline.Predict(pressure)

	return liquidVolume, vaporVolume, nil
}
