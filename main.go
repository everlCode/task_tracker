package main

import (
	"log"
	"os"
	"task-tracker/cli"
	"task-tracker/tasks"
	"task-tracker/tracker"
)

func main() {
	tasksFile, err := os.Open("tasks.json")
	if err != nil {
		log.Fatalln(err)
	}
	repository := tasks.New(tasksFile)
	tracker := tracker.New(repository)
	router := cli.New(tracker)

	router.Handle()
}
