package application

import (
	"github.com/paguerre3/as/internal/domain"
)

type PageData struct {
	RepairCode string
}

type DamagedSpaceshipUseCases interface {
	NextDamagedSystem() string
	RepairCode() *PageData
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
