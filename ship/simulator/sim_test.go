package main

import (
	"math"
	"testing"

	"github.com/353solutions/geo.v2"
)

func TestEuclideanBug(t *testing.T) {
	// Test the Euclidean function.
	d := geo.Euclidean(0, 0, 95e200, 168e200)
	if math.IsInf(d, 1) {
		t.Fatal(d)
	}
}
