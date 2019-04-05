package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/data", postData).Methods("POST")
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))

	result := r.URL.Query().Get("ok") == "ok"

	tmpl.Execute(w, struct{ Sucess bool }{result})
}

func postData(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	details := ContactDetails{
		Email:   email,
		Subject: subject,
		Message: message,
	}

	fmt.Println(details)

	http.Redirect(w, r, "/?ok=ok", 301)
}
