package main

import (
	"fmt"

	"github.com/lucasSSalgado/pokedex/internal/pokeapi"
)

func callbackMapB(s *stade) error {
	if s.Previous == nil {
		fmt.Println("not available url")
		return nil
	}
	locations, err := pokeapi.GetLocation(s.Previous)
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