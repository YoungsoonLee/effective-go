package main

import "fmt"

// Number is set of possible numbers.
type Number interface {
	~int | ~float64 // ~ mean. allow all integer and floating-point types
}

// Max retruns the maximal value in values.
func Max[T Number](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("Max of empty slice")
	}

	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}

	return max, nil
}

func main() {
	iVals := []int{1, 2, 3, 4, 5}
	fmt.Println(Max(iVals))

	fVals := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(Max(fVals))

	_, err := Max[int](nil)
	fmt.Println(err)
}
