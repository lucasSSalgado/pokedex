package main

import (
	"fmt"
	"bufio"
	"os"
)

type cliCommands struct {
	name string
	description string
}

func options() map[string] cliCommands {
	return map[string]cliCommands{
		"help": {
			name : "help",
			description: "show all available command and its usage",
		},
		"exit": {
			name : "exit",
			description: "exit the pokedex",
		},
	}
}

func evaluateInput(c string) {
	switch c {
		case "exit":
			os.Exit(0)
		case "help":
			helpCommand()
	}
}

func helpCommand() {
	sliceOfOption := options()

	for _, elem := range sliceOfOption {
		fmt.Println(elem.name + ": " +elem.description)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		evaluateInput(scanner.Text())
	}
}
