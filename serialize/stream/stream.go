package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Log is a log message.
type Log struct {
	Time    time.Time
	Level   int
	Message string
}

func ingestLogs(r io.Reader, handler func(Log)) error {

	dec := json.NewDecoder(r)

	for {
		var l Log
		err := dec.Decode(&l)
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		handler(l)
	}

	return nil
}

func printingHandler(l Log) {
	log.Printf("%#v\n", l)
}

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	go func() {
		defer w.Close()
		enc := json.NewEncoder(w)
		for i := 0; i < 5; i++ {
			l := Log{time.Now(), i, fmt.Sprintf("message #%d", i)}
			if err := enc.Encode(l); err != nil {
				log.Fatalf("error: %s", err)
			}
			time.Sleep(time.Second)
		}
	}()

	if err := ingestLogs(r, printingHandler); err != nil {
		log.Fatalf("error: %s", err)
	}

	json.NewEncoder(os.Stdout).Encode(Log{
		Time:    time.Now(),
		Level:   20,
		Message: "Database is up",
	})
}
