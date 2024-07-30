package main

import "fmt"

// Number is set of possible numbers.
type Number interface {
	~int | ~int16 | ~int32 | ~int64 |
		~uint | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Ring is a circular buffer.
type Ring[T Number] struct {
	size   int
	i      int
	values []T
}

// NewRing returns a new ring buffer with size.
func NewRing[T Number](size int) (*Ring[T], error) {
	if size <= 0 {
		return nil, fmt.Errorf("size must be > 0")
	}

	r := Ring[T]{
		size:   size,
		values: make([]T, size),
	}

	return &r, nil
}

// Push pushes item to the ring, possibly overwriting an old value.
func (r *Ring[T]) Push(value T) {
	r.values[r.i] = value
	r.i = (r.i + 1) % r.size
}

// Mean returns the mean of the ring values.
func (r *Ring[T]) Mean() float64 {
	var s T = 0
	for _, v := range r.values {
		s += v
	}

	return float64(s) / float64(r.size)
}

func main() {
	size := 4
	r, err := NewRing[int](size)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i <= 10; i++ {
		r.Push(i)
	}

	fmt.Println(r.Mean())
}
