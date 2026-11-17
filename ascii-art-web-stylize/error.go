package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func renderError(w http.ResponseWriter, status int) {
	page := "templates/" + strconv.Itoa(status) + ".html"

	tmpl, err := template.ParseFiles(page)
	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}

	w.WriteHeader(status)
	tmpl.Execute(w, nil)
}