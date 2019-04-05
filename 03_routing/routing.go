package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// gorilla follows the same url api signature
	subRouter := r.PathPrefix("/books").Subrouter()

	subRouter.HandleFunc("/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	subRouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "You've requested the book: %s\n", title)
	}).Methods("GET")
	http.ListenAndServe(":8080", r)
}
