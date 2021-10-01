package server

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed templates
var templateFS embed.FS

type templateRenderer struct {
	t *template.Template
}

func (r *templateRenderer) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	return r.t.ExecuteTemplate(w, name, data)
}

func (s *server) registerRenderer() {
	//s.http.Renderer = &templateRenderer{
	//	t: template.Must(template.ParseFS(templateFS, []string{
	//		"templates/*.html",
	//	}...)),
	//}
}
