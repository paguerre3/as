package domain

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

	Pc = 10.0 // Critical Pressure in MPa
	//Tc = 500.0     // Critical Temperature in Celsius
	Vc = 0.0035 // Critical Specific Volume in m^3/kg
	//Kl = 0.0002462 // Derived constant for liquid line
	//Kv = -3.014    // Derived constant for vapor line
	Kl = 0.00024623655913978494 // Corrected constant for liquid line
	Kv = -3.014623115577889     // Corrected constant for vapor line
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

func roundUp(num float64) float64 {
	const scale_min_digits = 100.0
	// Check if the number is effectively close to a whole number and needs rounding to two decimals
	if num == math.Floor(num) {
		// For whole numbers (like 30.0), round to exactly two decimal places
		return math.Round(num*scale_min_digits) / scale_min_digits
	}

	// Scale to 5 decimal places and round
	scale := 100000.0
	rounded := math.Round(num*scale) / scale
	if rounded == 0 {
		// double check this:
		return math.Round(rounded*scale_min_digits) / scale_min_digits
	}

	// Check if the number is very close to a whole number (e.g., 29.999 -> 30.00)
	str := fmt.Sprintf("%.2f", rounded)
	if rounded >= 1 && strings.HasSuffix(str, ".00") {
		num, _ := strconv.ParseFloat(str, 64)
		// double check this:
		rounded = math.Round(num*scale_min_digits) / scale_min_digits
	}

	// Return the rounded value with the needed precision
	return rounded
}

// Saturated liquid line equation
func (d *damagedSpaceshipImpl) SaturatedLiquidAndVaporVolumes(pressure float64) (float64, float64, error) {
	// Validate the pressure range
	//if pressure > Pc {
	//	return 0, 0, fmt.Errorf("pressure exceeds critical point (%f MPa)", Pc)
	//}
	pressureBD := decimal.NewFromFloat(pressure)
	PcBD := decimal.NewFromFloat(Pc)
	VcBD := decimal.NewFromFloat(Vc)
	KlBD := decimal.NewFromFloat(Kl)
	KvBD := decimal.NewFromFloat(Kv)

	// Calculate specific volumes
	//specificVolumeLiquidBD := Vc - Kl*(Pc-pressure)
	specificVolumeLiquidBD := VcBD.Sub(KlBD.Mul(PcBD.Sub(pressureBD)))
	//specificVolumeVaporBD := Vc + Kv*(pressure-Pc)
	specificVolumeVaporBD := VcBD.Add(KvBD.Mul(pressureBD.Sub(PcBD)))

	if specificVolumeLiquidBD.LessThan(decimal.Zero) || specificVolumeVaporBD.LessThan(decimal.Zero) {
		return 0, 0, fmt.Errorf("specific volumes cannot be negative")
	}

	svl, _ := specificVolumeLiquidBD.Float64()
	svv, _ := specificVolumeVaporBD.Float64()
	//return math.Max(svl, 0), math.Max(svv, 0), nil
	return roundUp(svl), roundUp(svv), nil
}
