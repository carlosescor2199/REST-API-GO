package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"

	Model "../model"
)

var Tasks = Model.Tasks

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Model.Task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		_, _ = fmt.Fprintf(w, "Insert a valid data")
	}

	_ = json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(Tasks) + 1

	Tasks = append(Tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	for _, task := range Tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask Model.Task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Pleasy enter valid data")
	}
	json.Unmarshal(reqBody, &updateTask)

	for index, task := range Tasks {
		if task.ID == taskID {
			Tasks = append(Tasks[:index], Tasks[index+1:]...)
			updateTask.ID = taskID
			Tasks = append(Tasks, updateTask)
			fmt.Fprintf(w, "Task %v deleted succesfully", taskID)
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	for index, task := range Tasks {
		if task.ID == taskID {
			Tasks = append(Tasks[:index], Tasks[index+1:]...)
			fmt.Fprintf(w, "Task %v updated succesfully", taskID)
		}
	}

}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}
