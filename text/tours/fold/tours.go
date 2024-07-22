package main

import (
	"fmt"
	"strings"
	"time"
)

// Tour is a struct that represents a tour.
type Tour struct {
	City string
	Name string
	Time time.Time
}

// findTour returns a tour by city and name.
func findTour(db []*Tour, city string) []*Tour {
	var tours []*Tour
	for _, t := range db {
		if strings.EqualFold(t.City, city) {
			tours = append(tours, t)
		}
	}

	return tours
}

// date is a shoorcut to create a time.Time.
func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func main() {
	db := []*Tour{
		{"Gdańsk", "Polish Food", date(2021, 1, 1)},
		{"Kraków", "Pub to Pub", date(2021, 1, 2)},
	}

	tours := findTour(db, "gdańsk")
	fmt.Printf("number of tours found: %d\n", len(tours))
}
