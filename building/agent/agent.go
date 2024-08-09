package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	version = "<unknown>"
)

const timeout = 3 * time.Second

func healthCheck(baseURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	url := fmt.Sprintf("%s/health", baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	return nil
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version and exit")
	flag.Parse()

	if showVersion {
		fmt.Printf("agent version %s\n", version)
	}
}
