package handler

import (
	// "ascii-art-web/handler"
	"ascii-art-web-stylize/ascii"
	"html/template"
	"net/http"
	"strings"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderError(w, http.StatusMethodNotAllowed, "400.html")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner == "" {
		RenderError(w, http.StatusBadRequest, "400.html")
		return
	}

	// 1. Convert carriage returns (\r\n) to plain newlines (\n)
	text = strings.ReplaceAll(text, "\r\n", "\n")
	// 2. Convert literal backslash+n ("\n") typed into input boxes into real line breaks
	text = strings.ReplaceAll(text, "\\n", "\n")

	artResult, err := ascii.Generate(text, banner)
	if err != nil {
	RenderError(w, http.StatusBadRequest, "400.html")
		return
	}

	tmpl, err := template.ParseFiles(getTemplatePath("index.html"))
	if err != nil {
	RenderError(w, http.StatusInternalServerError, "500.html")
		return
	}

	data := PageData{
		Result:    artResult,
		InputText: text,
		Banner:    banner,
	}

	if err := tmpl.Execute(w, data); err != nil {
		RenderError(w, http.StatusInternalServerError, "500.html")
		return
	}
}