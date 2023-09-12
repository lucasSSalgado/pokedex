package main

import (
	"github.com/lucasSSalgado/pokedex/internal/pokeapi"
)

type Stade struct {
	client pokeapi.Client
	pokemonList map[string]pokeapi.Pokemon
	Next     *string
	Previous *string	
}

func main() {
	stade := Stade{
		client: pokeapi.NewClient(),
		pokemonList: make(map[string]pokeapi.Pokemon),
	}
	startRelp(&stade)
}