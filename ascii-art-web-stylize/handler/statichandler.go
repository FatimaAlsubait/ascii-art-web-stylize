package handler

import (
	"os"
	"net/http"
	"path/filepath"
)
func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(r.URL.Path)
	filePath := filepath.Join(".", path)

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.IsDir() {
		RenderError(w, http.StatusNotFound, "404.html")
		return
	}

	http.ServeFile(w, r, filePath)
}