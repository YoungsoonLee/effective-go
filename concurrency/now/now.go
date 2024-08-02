package main

import (
	"sync/atomic"
	"time"
)

var (
	now atomic.Value
)

func init() {
	now.Store(time.Now())
	go func() {
		for {
			time.Sleep(time.Millisecond)
			now.Store(time.Now())
		}
	}()
}

// Now returns the current time.
func Now() time.Time {
	return now.Load().(time.Time)
}
