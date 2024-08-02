package main

import (
	"testing"
	"time"
)

var (
	t time.Time
)

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = time.Now()
	}
}

func BenchmarkAtomicTimeNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = Now()
	}
}
