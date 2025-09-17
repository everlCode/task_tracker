package main

import (
	"task-tracker/cli"
	"task-tracker/tracker"
)

func main() {
	tracker := tracker.New()
	router := cli.New(tracker)
	
	router.Handle()
}
