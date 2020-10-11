package model

type Task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []Task

var Tasks = allTasks{
	{
		ID:      1,
		Name:    "Task 1",
		Content: "This is task 1",
	},
}
