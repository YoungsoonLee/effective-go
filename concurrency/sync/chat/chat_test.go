package chat

import (
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type Sink struct {
	mu       sync.Mutex
	messages []string
}

func (s *Sink) Add(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.messages = append(s.messages, msg)
}

type Client struct {
	id   int
	sink *Sink
}

// Write implements io.Writer
func (c *Client) Write(p []byte) (n int, err error) {
	c.sink.Add(string(p))
	return len(p), nil
}

func TestNotify(t *testing.T) {
	var s Sink
	var clients []io.Writer
	const n = 3
	for i := 0; i < n; i++ {
		clients = append(clients, &Client{id: i, sink: &s})
	}

	r := Room{clients: clients}
	msg := "Who's on first?"
	r.Notify(msg)
	require.Equal(t, n, len(s.messages))
	for _, m := range s.messages {
		require.Equal(t, msg, m)
	}
}
