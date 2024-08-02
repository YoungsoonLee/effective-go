package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var (
	count int64
	m     sync.Mutex
)

func BenchmarkMutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Lock()
		count++
		m.Unlock()
	}
}

func BenchmarkAtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&count, 1)
	}
}
