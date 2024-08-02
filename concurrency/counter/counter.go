package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/exp/rand"
)

var (
	callCount int64
)

func handler(w http.ResponseWriter, r *http.Request) {
	//callCount++
	atomic.AddInt64(&callCount, 1)

	// Simulate work
	time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	w.Write([]byte("ok"))
}

func main() {
	const nRuns = 1000
	const nGoRoutines = 10

	var wg sync.WaitGroup
	wg.Add(nGoRoutines)

	for i := 0; i < nGoRoutines; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < nRuns; i++ {
				// Dummy ResponseWriter & http.Request for the handler
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
				handler(w, r)
			}
		}()
	}

	wg.Wait()
	println("callCount:", callCount)
}
