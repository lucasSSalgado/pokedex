package main

import "os"

func exit(s *Stade, args ...string) error {
	os.Exit(0)
	return nil
}