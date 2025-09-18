package tracker

import (
	"fmt"
	"log"
	"task-tracker/tasks"
)

type Tracker struct {
	repository *tasks.Repository
}

func New(repository *tasks.Repository) *Tracker {
	return &Tracker{
		repository: repository,
	}
}

func (tracker Tracker) Add(name string) {
	id, err := tracker.repository.Add(name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task added successfully (ID: %d)", id)
}

func (tracker Tracker) Update(id int, name string) {
	fmt.Println("Update METHOD EXECUTE")
}

func (tracker Tracker) Delete(id int) {
	fmt.Println("ADeleteDD METHOD EXECUTE")
}

func (tracker Tracker) MarkInProgress(id int) {
	fmt.Println("mark-in-progress METHOD EXECUTE")
}

func (tracker Tracker) MarkDone(id int) {
	fmt.Println("MarkDone METHOD EXECUTE")
}

func (tracker Tracker) List(id int) {
	fmt.Println("List METHOD EXECUTE")
}
