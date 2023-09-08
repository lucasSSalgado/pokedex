package main

type stade struct {
	Next *string
	Previous *string
}

func main() {
	stade := stade{
		Next: nil,
		Previous: nil,
	}

	startRelp(&stade)
}