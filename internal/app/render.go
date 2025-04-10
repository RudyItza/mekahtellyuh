package app

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (app *Application) Render(w http.ResponseWriter, r *http.Request, name string, data map[string]any) {
	tmplPath := filepath.Join("ui", "html", name)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		app.ServerError(w, err)
	}
}
