package main

import (
	"fmt"
	"os"
	"task-tracker/cli"
	"task-tracker/tasks"
	"task-tracker/tracker"
)

func main() {
	const tasksFilePath = "tasks.json"

	// Репозиторий теперь работает через путь, а не *os.File
	repository := tasks.New(tasksFilePath)
	tr := tracker.New(repository)
	router := cli.New(tr)

	if len(os.Args) < 2 {
		fmt.Println("Нужно указать команду")
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	router.Call(cmd, args)
}
