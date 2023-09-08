package main

import "os"

func exit(s *stade) error {
	os.Exit(0)
	return nil
}