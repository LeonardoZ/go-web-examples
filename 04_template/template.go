package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	http.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO List =)",
		Todos: []Todo{
			{Title: "Wake up", Done: true},
			{Title: "Shower", Done: true},
			{Title: "Eat", Done: false},
			{Title: "Go To workd", Done: false},
		},
	}
	tmpl := template.Must(template.ParseFiles("layout.html"))
	tmpl.Execute(w, data)

}
