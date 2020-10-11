package main

import (
	con "./controller"
	r "./router"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r.Router.HandleFunc("/", con.IndexRoute)
	r.Router.HandleFunc("/tasks", con.GetTasks).Methods("GET")
	r.Router.HandleFunc("/tasks", con.CreateTask).Methods("POST")
	r.Router.HandleFunc("/tasks/{id}", con.GetTask).Methods("GET")
	r.Router.HandleFunc("/tasks/{id}", con.DeleteTask).Methods("DELETE")
	r.Router.HandleFunc("/tasks/{id}", con.UpdateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", r.Router))
}
