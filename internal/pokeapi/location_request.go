package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(path *string) (LocationAreasResp, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if path != nil {
		url = *path	
	}
	
	// check the cache
	val, ok := c.memory.Get(url)
	if ok {
		fmt.Println("hits the cache!!!")
		location := LocationAreasResp{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationAreasResp{}, nil
		}
		return location, nil
	}

	// not in the cache yet
	fmt.Println("miss the cache!!!")
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	location := LocationAreasResp{}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// save in the cache
	c.memory.Add(url, bytes)
	
	return location, nil
}

func (c *Client) GetDetailsLocation(args ...string) (LocationDetails, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(args) < 2 {
		return LocationDetails{}, fmt.Errorf("location not provided")
	}

	fmt.Println(args)

	localName := args[1]
	fullPath := url + localName
	
	// check the cache
	val, ok := c.memory.Get(fullPath)
	if ok {
		fmt.Println("hits the cache!!!")
		location := LocationDetails{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return LocationDetails{}, nil
		}
		return location, nil
	}

	// not in the cache yet
	fmt.Println("miss the cache!!!")
	resp, err := http.Get(fullPath)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	location := LocationDetails{}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return LocationDetails{}, err
	}

	// save in the cache
	c.memory.Add(url, bytes)
	
	return location, nil
}