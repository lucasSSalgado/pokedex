package main

import (
	"fmt"
)

func callbackmap(s *Stade, args ...string) error {
	locations, err := s.client.GetLocation(s.Next)
	if err != nil {
		return err
	}	

	for _, elem := range locations.Results {
		fmt.Println(elem.Name)
	}
	
	s.Next = locations.Next
	s.Previous = locations.Previous
	return nil
}