package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

// Message is a generic interface for sending messages.
type Message struct {
	Text string
}

func processMessages(conn io.ReadWriteCloser, ch <-chan Message) error {
	dec := json.NewDecoder(conn)
	enc := json.NewEncoder(conn)

	for msg := range ch {
		if err := enc.Encode(msg); err != nil {
			return err
		}

		var reply struct {
			Output any
		}

		if err := dec.Decode(&reply); err != nil {
			return err
		}

		log.Printf("%#v -> %#v", msg, reply.Output)
	}

	return nil
}

var socketFile = "rpc.sock"

func main() {
	cmd := exec.Command(
		"go", "run", "server/main.go",
		"-socket", socketFile,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer cmd.Process.Kill()
	time.Sleep(time.Second)

	sock, err := net.Dial("unix", socketFile)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer sock.Close()

	ch := make(chan Message)
	go func() {
		ch <- Message{Text: "Was it a cat I saw?"}
	}()
}
