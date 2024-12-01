package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/as/internal/application"
)

type DamagedSpashipHandler interface {
	Status(c echo.Context) error
	Teapot(c echo.Context) error
}

type damagedSpaceshipHandler struct {
	damagedSpaceshipUseCases application.DamagedSpaceshipUseCases
}

func NewDamagedSpaceshipHandler(damagedSpaceshipUseCases application.DamagedSpaceshipUseCases) DamagedSpashipHandler {
	return &damagedSpaceshipHandler{
		damagedSpaceshipUseCases: damagedSpaceshipUseCases,
	}
}

func (d *damagedSpaceshipHandler) Status(c echo.Context) error {
	damagedSystem := d.damagedSpaceshipUseCases.NextDamagedSystem()
	return c.JSON(http.StatusOK, map[string]string{"damaged_system": damagedSystem})
}

func (d *damagedSpaceshipHandler) Teapot(c echo.Context) error {
	return c.NoContent(http.StatusTeapot)
}
