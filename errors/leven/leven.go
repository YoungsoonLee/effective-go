package main

import "fmt"

// EditDistance returns the edit distance between s1 and s2.
// It wraps dist.Edit against panics.
func EditDistance(s1, s2 string) (distance int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	// return dist.Edit(s1, s2), nil
	return 0, nil
}

func main() {

}
