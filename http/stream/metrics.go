package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

var (
	metrics   = []string{"CPU", "Memory", "Disk"}
	hosts     = []string{"srv1", "srv2", "srv3", "srv4", "srv5"}
	serverURL = "http://localhost:8080/metrics"
)

// Metric is a metric that can be collected by the metrics package.
type Metric struct {
	Name  string    `json:"name"`
	Host  string    `json:"host"`
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func collectMetrics() <-chan Metric {
	// This function collects metrics from the system and sends them to the metrics server.
	ch := make(chan Metric)
	go func() {
		for i := 0; i < 20; i++ {
			m := Metric{
				Name:  metrics[i%len(metrics)],
				Host:  hosts[i%len(hosts)],
				Time:  time.Now(),
				Value: rand.Float64() * 100,
			}
			ch <- m
			time.Sleep(117 * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}

func producer(w io.WriteCloser) {
	defer w.Close()

	enc := json.NewEncoder(w)
	for m := range collectMetrics() {
		if err := enc.Encode(m); err != nil {
			log.Printf("failed to encode: %s", err)
			return
		}
	}
}

func updateMetrics() error {
	// This function updates the metrics in the metrics server.
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}

	go producer(w)

	req, err := http.NewRequest("POST", serverURL, r)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func main() {
	if err := updateMetrics(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
