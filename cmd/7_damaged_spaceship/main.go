package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paguerre3/as/internal/api"
	"github.com/paguerre3/as/internal/application"
	"github.com/paguerre3/as/internal/common"
	"github.com/paguerre3/as/internal/infrastructure/web"
)

func registerWalletHandlers(e *echo.Echo, serverPort string) {

	e.Renderer = web.NewTemplateRenderer(common.TEMPLATES_DIR)

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // allow all origins
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// handlers
	uc := application.NewDamagedSpaceshipUseCases()
	api := api.NewDamagedSpaceshipHandler(uc)
	e.GET("/status", api.Status)
	e.POST("/teapot", api.Teapot)

	web := web.NewDamagedSpaceshipHandler(uc)
	e.GET("/repair-bay", web.RepairBay)
}

func main() {
	web.NewServerNode("Damaged-Spaceship", "0.0.0.0:8080", registerWalletHandlers).InitAndRun()
}
