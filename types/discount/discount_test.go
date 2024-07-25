package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// LineItem is an item in a shopping cart
type LineItem struct {
	Name   string
	Amount float64
	Price  float64
}

func Test_cartTotal(t *testing.T) {
	cart := []LineItem{
		{"lemon", 4, 0.5},
		{"orange", 5, 0.4},
		{"banana", 6, 0.1},
	}

	discounts := map[string]float64{
		"orange": 0.9, // 10% off
	}

	expected := 4*0.5 + 5*0.4*0.9 + 6*0.1
	total := cartTotal(cart, discounts)
	require.InDelta(t, expected, total, 0.001)
}

// cartTotal calculates the total price of a shopping cart
// discounts is a map from Name -> discount value.
func cartTotal(cart []LineItem, discounts map[string]float64) float64 {
	total := 0.0
	for _, li := range cart {
		// discount := discounts[li.Name]
		// total += li.Amount * li.Price * discount
		discount, ok := discounts[li.Name]
		if ok {
			total += li.Amount * li.Price * discount
		} else {
			total += li.Amount * li.Price
		}
	}
	return total
}
