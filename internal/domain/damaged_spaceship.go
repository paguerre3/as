package domain

import (
	"errors"
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

func (d *damagedSpaceshipImpl) SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error) {
	if pressure < PressureLine || pressure > PressureCritical {
		return 0, 0, fmt.Errorf("pressure %.2f MPa is out of range (%.2f MPa - %.2f MPa)", pressure, PressureLine, PressureCritical)
	}

	pressures := []float64{ReferencePressureMPa, CriticalPressureMPa}
	liquidVolumes := []float64{LiquidStartVolume, CriticalVolume}
	vaporVolumes := []float64{VaporStartVolume, CriticalVolume}

	liquidVolume, err := cubicSplineInterpolation(pressures, liquidVolumes, pressure)
	if err != nil {
		return 0, 0, err
	}

	vaporVolume, err := cubicSplineInterpolation(pressures, vaporVolumes, pressure)
	if err != nil {
		return 0, 0, err
	}

	lv, _ := decimal.NewFromFloat(liquidVolume).Round(5).Float64()
	gv, _ := decimal.NewFromFloat(vaporVolume).Round(5).Float64()

	return lv, gv, nil
}

// cubicSplineInterpolation implements cubic spline interpolation for a given set of data points.
func cubicSplineInterpolation(x, y []float64, xTarget float64) (float64, error) {
	n := len(x)
	if n != len(y) || n < 2 {
		return 0, errors.New("invalid input: x and y must have the same length and at least two points")
	}

	// Step 1: Calculate h and alpha
	h := make([]float64, n-1)
	alpha := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		h[i] = x[i+1] - x[i]
		if h[i] <= 0 {
			return 0, errors.New("x values must be strictly increasing")
		}
		alpha[i] = (y[i+1] - y[i]) / h[i]
	}

	// Step 2: Solve tridiagonal system for c
	l := make([]float64, n)
	mu := make([]float64, n)
	z := make([]float64, n)

	l[0], l[n-1] = 1, 1
	for i := 1; i < n-1; i++ {
		l[i] = 2*(x[i+1]-x[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}

	c := make([]float64, n)
	for j := n - 2; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
	}

	// Step 3: Calculate b and d
	b := make([]float64, n-1)
	d := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = (y[i+1]-y[i])/h[i] - h[i]*(c[i+1]+2*c[i])/3
		d[i] = (c[i+1] - c[i]) / (3 * h[i])
	}

	// Step 4: Interpolate
	for i := 0; i < n-1; i++ {
		if xTarget >= x[i] && xTarget <= x[i+1] {
			diff := xTarget - x[i]
			return y[i] + b[i]*diff + c[i]*diff*diff + d[i]*diff*diff*diff, nil
		}
	}

	return 0, errors.New("xTarget out of range")
}
