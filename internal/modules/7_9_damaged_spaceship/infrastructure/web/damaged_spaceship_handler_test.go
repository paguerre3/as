package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/as/internal/modules/7_9_damaged_spaceship/application"
	common_infra "github.com/paguerre3/as/internal/modules/common/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestDamagedSpaceshipRepairBay(t *testing.T) {
	// Create a test Echo instance
	e := echo.New()
	e.Renderer = NewTemplateRenderer(common_infra.TEMPLATES_DIR)

	uc := application.NewDamagedSpaceshipUseCases()
	// Create a test WalletHandler instance
	webHandler := NewDamagedSpaceshipHandler(uc)

	// Test case 1: Index method
	req, err := http.NewRequest("GET", "/repair-bay", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = webHandler.RepairBay(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Header().Get("Content-Type"), "text/html")
	strBody := rec.Body.String()
	assert.Contains(t, strBody, "<!DOCTYPE html>")
	assert.Contains(t, strBody, "<title>Repair</title>")
}
