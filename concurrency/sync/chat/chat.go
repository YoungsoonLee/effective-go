package chat

import (
	"io"
	"sync"
)

// Room is a chat room
type Room struct {
	clients []io.Writer
}

// Notify sends msg to all clients in parallel.
// It will return after all messages are sent.
func (r *Room) Notify(msg string) {
	var done sync.WaitGroup
	done.Add(len(r.clients))
	for _, c := range r.clients {
		go func(client io.Writer) {
			defer done.Done()
			client.Write([]byte(msg))
		}(c)
	}
	done.Wait()
}
