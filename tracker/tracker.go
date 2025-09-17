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
