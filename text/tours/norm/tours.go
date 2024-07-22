package main

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/unicode/norm"
)

// Tour is a struct that represents a tour.
type Tour struct {
	City string
	Name string
	Time time.Time
}

// normString normalizes a string in NFKC.
func normString(s string) string {
	return norm.NFKC.String(s)
}

// NewTour returns a new tour.
func NewTour(city, name string, time time.Time) *Tour {
	return &Tour{
		City: normString(city),
		Name: name,
		Time: time,
	}
}

// findTour returns all tours by city.
func findTour(db []*Tour, city string) []*Tour {
	city = normString(city)
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
