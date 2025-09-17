package tasks

import "os"

type Repository struct {
	tasks []Task
}

type Task struct {
	Id     int
	Name   string
	Status int
}

func New(file *os.File) *Repository {
	var tasks []Task = getTasksFromFile(file)

	return &Repository{
		tasks: tasks,
	}
}

func getTasksFromFile(file *os.File) []Task {
	return []Task{}
}
