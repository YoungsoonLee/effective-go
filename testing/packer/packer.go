package packer

import "fmt"

// Item is an package item.
type Item struct {
	Name   string
	Weight float64
}

// Pack returns items split to boxes.
// It considers item weight only and uses a greedy algorithm.
func Pack(boxCap float64, cart []Item) ([][]Item, error) {
	var boxes [][]Item
	weight := 0.0
	var box []Item
	for _, item := range cart {
		if item.Weight > boxCap {
			const format = "item %#v is too heavy for %f"
			return nil, fmt.Errorf(format, item, boxCap)
		}

		if weight+item.Weight > boxCap { /// need a new box
			boxes = append(boxes, box)
			box = []Item{}
			weight = 0.0
		}

		box = append(box, item)
		weight += item.Weight
	}

	if len(box) > 0 { // add sentinel
		boxes = append(boxes, box)
	}

	return boxes, nil
}
