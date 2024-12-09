package main

import (
	"bytes"
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	damaged_spaceship_api "github.com/paguerre3/as/internal/modules/7_9_damaged_spaceship/api"
	damaged_spaceship_app "github.com/paguerre3/as/internal/modules/7_9_damaged_spaceship/application"
	damaged_spaceship_web "github.com/paguerre3/as/internal/modules/7_9_damaged_spaceship/infrastructure/web"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
)

func logRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		log.Printf("Incoming request %s", strings.Repeat("-", 100))
		// Log HTTP method (GET, POST, etc.)
		log.Printf("Method: %s", req.Method)

		// Log URI (the full URL path)
		log.Printf("URI: %s", req.RequestURI)

		// Log headers
		log.Printf("Headers: %v", req.Header)

		// Log query parameters (if any)
		queryParams := req.URL.Query()
		log.Printf("Query Parameters: %v", queryParams)

		// Log remote address (client IP)
		log.Printf("Remote Address: %s", req.RemoteAddr)

		// Log the host (server host)
		log.Printf("Host: %s", req.Host)

		// Log the referer (if present)
		log.Printf("Referer: %s", req.Referer())

		// Log user agent (browser or client)
		log.Printf("User-Agent: %s", req.UserAgent())

		// Read and log the body
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Failed to read body: %v", err)
			return err
		}

		// Log the raw body content
		log.Printf("Body: %s", string(bodyBytes))

		// Restore the body so it can be read by the next handler
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		return next(c)
	}
}

func registerWalletHandlers(e *echo.Echo, serverPort string) {

	e.Renderer = damaged_spaceship_web.NewTemplateRenderer(common_infra.TEMPLATES_DIR)

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // allow all origins
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Apply middleware to log raw requests
	e.Use(middleware.Logger())
	e.Use(logRequest)

	// handlers
	uc := damaged_spaceship_app.NewDamagedSpaceshipUseCases()
	api := damaged_spaceship_api.NewDamagedSpaceshipHandler(uc)
	e.GET("/status", api.Status)
	e.POST("/teapot", api.Teapot)
	web := damaged_spaceship_web.NewDamagedSpaceshipHandler(uc)
	e.GET("/repair-bay", web.RepairBay)

	e.GET("/phase-change-diagram", api.PhaseChangeDiagram)
}

func main() {
	damaged_spaceship_web.NewServerNode("Damaged-Spaceship-Server", "0.0.0.0:8080", registerWalletHandlers).InitAndRun()
}
