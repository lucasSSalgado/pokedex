package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetLocation(path *string) (LocationAreasResp, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if path != nil {
		url = *path	
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	bytes, err :=io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	location := LocationAreasResp{}
	err = json.Unmarshal(bytes, &location)
	if err != nil {
		return LocationAreasResp{}, err
	}

	
	return location, nil
}