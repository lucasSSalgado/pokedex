package main

import "fmt"

func callbackInspect(s *Stade, args ...string) error {
	allPookemons := s.pokemonList

	if len(args) < 2 {
		return fmt.Errorf("pokemon name not provided")
	}
	pokemom, ok := allPookemons[args[1]]
	if !ok {
		return fmt.Errorf("pokemon not found in your pokedex")
	}

	fmt.Printf("Name: %s\n", pokemom.Name)
	fmt.Printf("Weight: %d\n", pokemom.Weight)
	fmt.Printf("Height: %d\n", pokemom.Height)
	fmt.Println("Stats: ")
	for _, val := range pokemom.Stats {
		fmt.Printf("- %s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Abilities: ")
	for _, val := range pokemom.Abilities {
		fmt.Printf("- %s\n", val.Ability.Name)
	}

	return nil
}