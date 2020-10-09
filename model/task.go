package model

type Task struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	Content string `json:"Content"`
}
