// MY FIRST BACKEND TO A SERVER IN GOLANG 
// server was created for custom made CTF game
package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	student := Student{
		Name: "ctf player",
	}
	parsedTemplate, _ := template.ParseFiles("static/static.html")
	err := parsedTemplate.Execute(w, student)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}

func main() {
	r := mux.NewRouter() // Routes consist of a path and a handler function.
	r.HandleFunc("/", renderTemplate).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
