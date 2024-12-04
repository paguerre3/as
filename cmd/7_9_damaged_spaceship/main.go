package main

import (
	"bytes"
	"io"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paguerre3/as/internal/api"
	"github.com/paguerre3/as/internal/application"
	"github.com/paguerre3/as/internal/common"
	"github.com/paguerre3/as/internal/infrastructure/web"
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

	e.Renderer = web.NewTemplateRenderer(common.TEMPLATES_DIR)

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // allow all origins
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Apply middleware to log raw requests
	e.Use(middleware.Logger())
	e.Use(logRequest)

	// handlers
	uc := application.NewDamagedSpaceshipUseCases()
	api := api.NewDamagedSpaceshipHandler(uc)
	e.GET("/status", api.Status)
	e.POST("/teapot", api.Teapot)
	web := web.NewDamagedSpaceshipHandler(uc)
	e.GET("/repair-bay", web.RepairBay)

	e.GET("/phase-change-diagram", api.PhaseChangeDiagram)
}

func main() {
	web.NewServerNode("Damaged-Spaceship", "0.0.0.0:8080", registerWalletHandlers).InitAndRun()
}
