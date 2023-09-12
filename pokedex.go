package main

import "fmt"

func callbackPokedex(s *Stade, args ...string) error {
	all := s.pokemonList
	fmt.Println("Your pokedex: ")

	if len(all) == 0 {
		fmt.Println("emoty pokedex")
		return nil
	}
	for _, val := range all {
		fmt.Printf("- %s\n", val.Name)
	}
	return nil
}