package html

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *
var files embed.FS

var (
	dashboard = parse("dashboard.html")
)

type DashboardParams struct {
	Title   string
	Message string
}

func Dashboard(w io.Writer, p DashboardParams) error {
	return dashboard.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}
