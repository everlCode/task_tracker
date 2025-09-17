package tracker

import "fmt"

type Tracker struct {
}

func New() *Tracker {
	return &Tracker{}
}

func (tracker Tracker) Add(name string) {
	fmt.Println("ADD METHOD EXECUTE")
	fmt.Println(name)
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
