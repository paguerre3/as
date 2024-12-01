package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ServerNode interface {
	InitAndRun()
}

type serverNode struct {
	name             string
	serverPort       string
	registerHandlers func(e *echo.Echo, serverPort string)
}

func NewServerNode(name, serverPort string,
	registerHandlers func(e *echo.Echo, serverPort string)) ServerNode {
	return &serverNode{
		name:             name,
		serverPort:       serverPort,
		registerHandlers: registerHandlers,
	}
}

func (s *serverNode) InitAndRun() {
	log.Infof("Starting %s Server on TCP Port %s", s.name, s.serverPort)
	// Initialize a new Echo instance
	e := echo.New()

	if s.registerHandlers == nil {
		panic("registerHandlers is nil")
	}

	// Pass echo server instance and port
	s.registerHandlers(e, s.serverPort)

	// Start the server on the specified port
	e.Start(s.serverPort)
}
