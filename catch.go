package main

import (
	"fmt"
	"math/rand"
)

func callbackcatch(s *Stade, args ...string) error {
	pokemon, err := s.client.CatchPokemon(args...)
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience
	playerLocky := rand.Intn(baseExp)
	basePowerNeeded := 50

	if basePowerNeeded > playerLocky {
		s.pokemonList[pokemon.Name] = pokemon
		fmt.Printf("%s was catch\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s scape, try again\n", pokemon.Name)
	return nil
}