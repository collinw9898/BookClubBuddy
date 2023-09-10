package web

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("web/templates", "layout.html")
	fp := filepath.Join("web/templates", "index.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "layout", nil)
}
