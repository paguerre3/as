package application

import (
	domain "github.com/paguerre3/as/internal/modules/7_9_damaged_spaceship/domain"
)

type PageData struct {
	RepairCode string
}

type Response struct {
	SpecificVolumeLiquid float64 `json:"specific_volume_liquid"`
	SpecificVolumeVapor  float64 `json:"specific_volume_vapor"`
}

type DamagedSpaceshipUseCases interface {
	NextDamagedSystem() string
	RepairCode() *PageData
	PhaseChangeDiagram(pressure float64) (*Response, error)
}

type damagedSpaceshipUseCasesImpl struct {
	damagedSpaceship domain.DamagedSpaceship
}

func NewDamagedSpaceshipUseCases() DamagedSpaceshipUseCases {
	return &damagedSpaceshipUseCasesImpl{
		damagedSpaceship: domain.NewDamagedSpaceship(),
	}
}

func (d *damagedSpaceshipUseCasesImpl) NextDamagedSystem() string {
	return d.damagedSpaceship.NextDamagedSystem()
}

func (d *damagedSpaceshipUseCasesImpl) RepairCode() *PageData {
	repairCode, ok := d.damagedSpaceship.RepairCode()
	if !ok {
		repairCode = domain.PICK_SYSTEM
	}
	return &PageData{
		RepairCode: repairCode,
	}
}

func (d *damagedSpaceshipUseCasesImpl) PhaseChangeDiagram(pressure float64) (*Response, error) {
	vLiquid, vVapor, err := d.damagedSpaceship.SaturatedLiquidAndVaporVolumes(pressure)
	if err != nil {
		return nil, err
	}
	return &Response{
		SpecificVolumeLiquid: vLiquid,
		SpecificVolumeVapor:  vVapor,
	}, nil
}
