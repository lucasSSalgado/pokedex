package main

import "fmt"

func help(s *stade) error {
	availableCommands := getCommands()
	fmt.Println("Welcome to pokedex, here is a list of available commands: ")
	for _, elem := range availableCommands {
		fmt.Println("- " + elem.name + ": " + elem.description)
	}
	return nil
}