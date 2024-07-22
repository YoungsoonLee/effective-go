package main

import "log"

// LocationEvent is a struct that represents a location event.
type LocationEvent struct {
	ID    string
	Lati  float64
	Longi float64
}

func main() {
	// Create a new LocationEvent.
	evt := &LocationEvent{
		ID:    "123",
		Lati:  37.7749,
		Longi: -122.4194,
	}

	// Print the event.
	// # is type
	log.Printf("info: loc: %#v", evt)
}
