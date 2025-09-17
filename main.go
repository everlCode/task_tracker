package main

import (
	"fmt"
	"os"
)

func main() {
	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	} else {
		fmt.Println("No command provided")
		return
	}

	fmt.Println(command)
}
