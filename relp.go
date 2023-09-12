package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommands struct {
	name string
	description string
	callback func(s *Stade, args ...string) error
}

func startRelp(s *Stade) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")
		scanner.Scan()

		userTexted := cleanInputs(scanner.Text())
		if len(userTexted) == 0 {
			continue
		}

		command := userTexted[0]
		options := getCommands()

		function, ok := options[command]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := function.callback(s, userTexted...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanInputs(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}

func getCommands()map[string]cliCommands {
	return map[string]cliCommands{
		"exit": {
			name: "exit",
			description: "kill the program",
			callback: exit,
		},
		"help": {
			name: "help",
			description: "show all available command and its usage",
			callback: help,
		},
		"map": {
			name: "map",
			description: "get 20 locations",
			callback: callbackmap,
		},
		"mapb": {
			name: "mapb",
			description: "get the last 20 location if possible",
			callback: callbackMapB,
		},
		"explore": {
			name: "explore",
			description: "get pokemons available in the {location}",
			callback: callbackExplore,
		},
		"catch": {
			name: "catch",
			description: "try to catch the pokemon base in the {name}",
			callback: callbackcatch,
		},
		"inspect": {
			name: "inspect",
			description: "check the status of a pokemon if early catch {name}",
			callback: callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "return all available pokemons",
			callback: callbackPokedex,
		},
	}
}