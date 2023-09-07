package main

import (
		"bufio"
		"fmt"
		"log"
		"os"

		"github.com/lucasSSalgado/pokedex/pokeapi"
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
		"map": {
			name: "map",
			description: "show the next 20 location available",
		}, 
		"mapb": {
			name: "mapb",
			description: "show the previos 20 location available",
		},
	}
}

func evaluateInput(c string) {
	switch c {
		case "exit":
			os.Exit(0)
		case "help":
			helpCommand()
		case "map":
			mapCommand(&pokeapi.InitialStruct ,true)
		case "mapb":
			mapCommand(&pokeapi.InitialStruct ,false)
	}
}

func helpCommand() {
	sliceOfOption := options()
	for _, elem := range sliceOfOption {
		fmt.Println(elem.name + ": " + elem.description)
	}
}

func mapCommand(json *pokeapi.LocationAreaJson, i bool) {
	list, err := pokeapi.CallLocation(json, i)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, elem := range list {
		fmt.Println(elem)
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