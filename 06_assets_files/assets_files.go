package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets/"))

	// different names
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
