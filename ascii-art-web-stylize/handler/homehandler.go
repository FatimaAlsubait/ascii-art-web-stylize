package handler

import (
	"html/template"
	"net/http"


)

type PageData struct {
	Result    string
	InputText string
	Banner    string
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound, "404.html")
		return
	}

	tmpl, err := template.ParseFiles(getTemplatePath("index.html"))
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "500.html")
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		RenderError(w, http.StatusInternalServerError, "500.html")
		return
	}
}