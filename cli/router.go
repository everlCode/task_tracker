package cli

import (
	"fmt"
	"reflect"
	"strconv"
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

func (r Router) Call(command string, args []string) {
	// Делаем первую букву заглавной, чтобы совпадало с экспортируемым методом
	command = strings.ToUpper(string(command[0])) + command[1:]
	commandSlice := strings.Split(command, "-")
	for k, word := range commandSlice {
		word = strings.ToUpper(string(word[0])) + word[1:]
		commandSlice[k] = word
	}

	command = strings.Join(commandSlice, "")

	reflection := reflect.ValueOf(r.Tracker)
	method := reflection.MethodByName(command)

	if !method.IsValid() {
		fmt.Println("Метод не найден")
		return
	}

	methodType := method.Type()

	if len(args) != methodType.NumIn() {
		fmt.Printf("Ожидается %d аргументов, передано %d\n", methodType.NumIn(), len(args))
		return
	}

	callArgs := []reflect.Value{}
	for i := 0; i < methodType.NumIn(); i++ {
		paramType := methodType.In(i)
		rawArg := args[i]

		var val reflect.Value
		switch paramType.Kind() {
		case reflect.String:
			val = reflect.ValueOf(rawArg)
		case reflect.Int:
			parsed, err := strconv.Atoi(rawArg)
			if err != nil {
				fmt.Printf("Ошибка преобразования '%s' в int\n", rawArg)
				return
			}
			val = reflect.ValueOf(parsed)
		case reflect.Bool:
			parsed, err := strconv.ParseBool(rawArg)
			if err != nil {
				fmt.Printf("Ошибка преобразования '%s' в bool\n", rawArg)
				return
			}
			val = reflect.ValueOf(parsed)
		default:
			fmt.Printf("Тип %s не поддерживается\n", paramType.Kind())
			return
		}

		callArgs = append(callArgs, val)
	}

	method.Call(callArgs)
}
