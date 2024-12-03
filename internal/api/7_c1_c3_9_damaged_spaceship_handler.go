package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/as/internal/application"
)

type DamagedSpashipHandler interface {
	Status(c echo.Context) error
	Teapot(c echo.Context) error
	PhaseChangeDiagram(c echo.Context) error
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

func (d *damagedSpaceshipHandler) PhaseChangeDiagram(c echo.Context) error {
	pressureParam := c.QueryParam("pressure")
	if pressureParam == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing 'pressure' parameter"})
	}
	// Convert "pressure" to a float
	pressure, err := strconv.ParseFloat(pressureParam, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid 'pressure' parameter"})
	}
	response, err := d.damagedSpaceshipUseCases.PhaseChangeDiagram(pressure)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, response)
}
