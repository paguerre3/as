package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/as/internal/application"
)

type DamagedSpashipHandler interface {
	RepairBay(c echo.Context) error
}

type damagedSpaceshipHandler struct {
	damagedSpaceshipUseCases application.DamagedSpaceshipUseCases
}

func NewDamagedSpaceshipHandler(damagedSpaceshipUseCases application.DamagedSpaceshipUseCases) DamagedSpashipHandler {
	return &damagedSpaceshipHandler{
		damagedSpaceshipUseCases: damagedSpaceshipUseCases,
	}
}

func (d *damagedSpaceshipHandler) RepairBay(c echo.Context) error {
	return c.Render(http.StatusOK, "repare-bay.html", d.damagedSpaceshipUseCases.RepairCode())
}
