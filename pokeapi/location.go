package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreaJson struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var InitialStruct LocationAreaJson
func init() {
	firstCall := "https://pokeapi.co/api/v2/location-area/"
	InitialStruct = LocationAreaJson{
		Next: &firstCall,
	}
}

func CallLocation(l *LocationAreaJson, direction bool) ([]string, error) {
	if l.Previous == nil && !direction{
		fmt.Println("no previus location found")
		return nil, nil
	}

	var url string
	if direction {
		url = *l.Next
	} else {
		url = *l.Previous
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = json.Unmarshal(body, &l)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	listOfLocation := make([]string, 0)
	for _, elem := range l.Results {
		listOfLocation = append(listOfLocation, elem.Name)
	}

	return listOfLocation, nil
}