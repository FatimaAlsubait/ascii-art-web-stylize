package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)


func getTemplatePath(page string) string {
	return filepath.Join("templates", page)
}
func RenderError(w http.ResponseWriter, status int, page string) {
	tmplPath := getTemplatePath(page)
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := tmpl.Execute(w, nil); err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "500 Internal Server Error")
		return
	}
}