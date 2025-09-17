package cli

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"task-tracker/tracker"
)

type Router struct {
	Tracker tracker.Tracker
}

func New(tracker *tracker.Tracker) *Router {
	router := &Router{
		Tracker: *tracker,
	}

	return router
}

func (r Router) Handle() {
	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		fmt.Println("No command provided")
		return
	}

	r.isCommandExist(command)

	// if _, ok := t.MethodByName("Goodbye"); ok {
	// 	fmt.Println("Метод найден")
	// } else {
	// 	fmt.Println("Метод не найден")
	// }

	// fmt.Println(command)
}

func (r Router) isCommandExist(command string) {
	command = strings.ToUpper(string(command[0])) + command[1:]
	reflection := reflect.ValueOf(r.Tracker)

	method := reflection.MethodByName(command)

	if !method.IsValid() {
		fmt.Println("Метод не найден")
		return
	} else {
		arg := []reflect.Value{
			reflect.ValueOf("Buy chocolade"),
		}
		method.Call(arg)
	}
}
