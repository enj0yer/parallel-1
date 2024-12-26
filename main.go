package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("no command provided\n")
		fmt.Println("try command list to show all possible commands")
		return
	}
	name := os.Args[1]
	command, ok := Commands[name]
	if !ok {
		fmt.Printf("command with name %s not found\n", name)
		return
	}
	var commandArgs []string
	if len(os.Args) == 2 {
		commandArgs = make([]string, 0)
	} else {
		commandArgs = os.Args[2:]
	}
	if output, err := command.Run(commandArgs); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	} else {
		fmt.Println(output)
	}

}
