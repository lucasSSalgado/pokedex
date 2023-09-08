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
	callback func(s *stade) error
}

func startRelp(s *stade) {
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
		err := function.callback(s)
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
	}
}