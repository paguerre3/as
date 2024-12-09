package web

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

var (
	templatesPathResolver = NewPathResolver()
)

type TemplateRenderer interface {
	Render(w io.Writer, name string, data any, c echo.Context) error
}

// TemplateRenderer is a custom Echo template renderer
type templateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(templatesPathOfDomain string) TemplateRenderer {
	templateDir := templatesPathResolver(templatesPathOfDomain)
	templatesPath := fmt.Sprintf("%s/*.html", templateDir)
	return &templateRenderer{
		templates: template.Must(template.ParseGlob(templatesPath)),
	}
}

func (t *templateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
