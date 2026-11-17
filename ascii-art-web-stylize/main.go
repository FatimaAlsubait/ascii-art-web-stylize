package main

import (
	"ascii-art-web-stylize/handler"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/ascii-art", handler.AsciiArtHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
