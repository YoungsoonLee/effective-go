package main

import (
	"fmt"
	"log"
)

// Server is an HTTP server
type Server struct {
	verbose bool
	port    int
}

// NewServer returns a Server with options.
func NewServer(options ...func(*Server) error) (*Server, error) {
	srv := &Server{
		port: 8090, // default port
	}

	for _, option := range options {
		if err := option(srv); err != nil {
			return nil, err
		}
	}

	return srv, nil
}

// WithVerbose sets the verbose option.
func WithVerbose(s *Server) error {
	s.verbose = true
	return nil
}

const portErrFmt = "port must be between 0 and %d, got %d"

// WithPort sets the port option.
func WithPort(port int) func(*Server) error {
	const maxPort = 0xFFFF
	return func(s *Server) error {
		if port <= 0 || port > maxPort {
			return fmt.Errorf(portErrFmt, maxPort, port)
		}
		s.port = port
		return nil
	}
}

func main() {
	srv, err := NewServer(
		WithPort(8888),
	)

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	WithVerbose(srv)

	fmt.Printf("%#v\n", srv)
}
