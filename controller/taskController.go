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

type Task struct {
	Model.Task
}

type allTasks []Task

var tasks = allTasks {
}

func (this *Task) getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func (this *Task) createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid data")
	}

	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func (this *Task) getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func (this *Task) updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask Task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Pleasy enter valid data")
	}
	json.Unmarshal(reqBody, &updateTask)

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index + 1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)
			fmt.Fprintf(w, "Task %v deleted succesfully", taskID)
		}
	}
}

func (this *Task) deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	for index, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:index], tasks[index + 1:]...)
			fmt.Fprintf(w, "Task %v updated succesfully", taskID)
		}
	}

}

func main() {
	task := new(Task)
	task.ID=(len(tasks)+1)
	task.Name="Task 1"
	task.Content="Some task"
	tasks = append(tasks, *task)
}