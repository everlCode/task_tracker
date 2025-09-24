package tracker

import (
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
	id, err := tracker.repository.Update(id, name)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task updated successfully (ID: %d)", id)
}

func (tracker Tracker) Delete(id int) {
	id, err := tracker.repository.Delete(id)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task deleted successfully (ID: %d)", id)
}

func (tracker Tracker) MarkInProgress(id int) {
	id, err := tracker.repository.UpdateStatus(id, tasks.StatusInProgress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task marked in progress successfully (ID: %d)", id)
}

func (tracker Tracker) MarkDone(id int) {
	id, err := tracker.repository.UpdateStatus(id, tasks.StatusDone)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task marked done successfully (ID: %d)", id)
}

func (tracker Tracker) List(id int) {
	id, err := tracker.repository.UpdateStatus(id, tasks.StatusInProgress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Task deleted successfully (ID: %d)", id)
}
