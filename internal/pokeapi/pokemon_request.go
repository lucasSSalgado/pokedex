package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(args ...string) (Pokemon, error) {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"

	if len(args) < 2 {
		return Pokemon{}, fmt.Errorf("pokemon name not provided!")
	}

	pokeName := args[1]
	fullPath := baseUrl + pokeName

	req, err := http.NewRequest("GET", fullPath, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(bytes, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}