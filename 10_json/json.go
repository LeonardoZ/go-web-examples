package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lasttname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", Decode)
	http.HandleFunc("/encode", Encode)
	http.ListenAndServe(":8080", nil)
}

func Decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func Encode(w http.ResponseWriter, r *http.Request) {
	leo := User{
		Firstname: "Leonardo",
		Lastname:  "Zapparoli",
		Age:       26}
	json.NewEncoder(w).Encode(leo)
}
