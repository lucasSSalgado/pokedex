package main

import (
	"fmt"

	"github.com/lucasSSalgado/pokedex/internal/pokeapi"
)

func callbackmap(s *stade) error {
	locations, err := pokeapi.GetLocation(s.Next)
	if err != nil {
		return err
	}

	s.Next = locations.Next
	s.Previous = locations.Previous

	for _, elem := range locations.Results {
		fmt.Println(elem.Name)
	}
	return nil
}