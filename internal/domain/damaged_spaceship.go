package domain

import (
	"fmt"
	"sync"
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

	// Known data points for interpolation
	dataPoints = []struct {
		pressure float64
		vLiquid  float64
		vVapor   float64
	}{
		{1, 0.0010, 0.0200},
		{2, 0.00105, 0.0180},
		{3, 0.0011, 0.0150},
		{4, 0.0012, 0.0120},
		{5, 0.00125, 0.0090},
		{6, 0.0013, 0.0070},
		{7, 0.0014, 0.0050},
		{8, 0.0016, 0.0040},
		{9, 0.002, 0.0036},
		{10, 0.0035, 0.0035}, // Critical point
		{11, 0.0036, 0.0034},
		{12, 0.00365, 0.0033},
		{13, 0.0037, 0.0032},
		{14, 0.00375, 0.0031},
		{15, 0.0038, 0.0030},
		{16, 0.00385, 0.0029},
		{17, 0.0039, 0.0028},
		{18, 0.00395, 0.0027},
		{19, 0.004, 0.0026},
		{20, 0.00405, 0.0025},
		{25, 0.00425, 0.0020},
		{30, 0.0045, 0.0018},
		{35, 0.00475, 0.0016},
		{40, 0.005, 0.0015},
		{45, 0.00525, 0.0014},
		{50, 0.0055, 0.0013},
	}
)

const (
	PICK_SYSTEM = "<pick one of the systems>"

	P_c = 10.0   // MPa
	v_c = 0.0035 // m^3/kg
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

// Linear interpolation function
func (d *damagedSpaceshipImpl) interpolate(pressure, p1, p2, v1, v2 float64) float64 {
	return v1 + (pressure-p1)*(v2-v1)/(p2-p1)
}

// Saturated liquid line equation
func (d *damagedSpaceshipImpl) SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error) {
	if pressure < 0.05 || pressure > 50.0 {
		return 0, 0, fmt.Errorf("pressure out of range")
	}

	var vLiquid, vVapor float64
	for i := 0; i < len(dataPoints)-1; i++ {
		p1 := dataPoints[i].pressure
		p2 := dataPoints[i+1].pressure
		if pressure >= p1 && pressure <= p2 {
			vLiquid = d.interpolate(pressure, p1, p2, dataPoints[i].vLiquid, dataPoints[i+1].vLiquid)
			vVapor = d.interpolate(pressure, p1, p2, dataPoints[i].vVapor, dataPoints[i+1].vVapor)
			break
		}
	}
	return vLiquid, vVapor, nil
}
