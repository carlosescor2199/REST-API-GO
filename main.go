package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	con "./controller"
)


func indexRoute(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	task := new(con.Task)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", congetTasks).Methods("GET")
	router.HandleFunc("/tasks", controller.createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", contacto.getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", contacto.deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":4000", router))
}
