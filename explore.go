package main

import (
	"fmt"
)

func callbackExplore(s *Stade, args ...string) error {
	details, err := s.client.GetDetailsLocation(args...)
	if err != nil {
		return err
	}

	fmt.Println("Available pokemons: ")

	for _, val := range details.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}

	return nil
}