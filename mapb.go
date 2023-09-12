package main

import (
	"fmt"
)

func callbackMapB(s *Stade, args ...string) error {
	if s.Previous == nil {
		fmt.Println("not available url")
		return nil
	}
	locations, err := s.client.GetLocation(s.Previous)
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