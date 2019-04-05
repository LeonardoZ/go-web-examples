package main

import (
	"fmt"
	"net/http"
)

func main() {
	// dynamic requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "User"
		}
		fmt.Fprintf(w, "Welcome to my website, %s\n!", name)
	})

	// static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Accepting connections
	http.ListenAndServe(":8080", nil)
}
