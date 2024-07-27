package main

import "testing"

var (
	values   []int
	size     = 9323
	expected int
)

func init() {
	for i := 0; i < size; i++ {
		values = append(values, i+1)
	}
	expected = size * ((values[0] + values[size-1]) / 2)
}

func BenchmarkCumsum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := cumsum(values)
		if len(out) != len(values) {
			b.Fatalf("expected %d, got %d", len(values), len(out))
		}
	}
}

// go test -bench . -cpuprofile cpu.pprof
// go tool pprof cpu.pprof
