package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func totalPayments(r io.Reader) (float64, error) {

	var reply struct {
		Payments []struct {
			Amount float64
		}
	}

	if err := json.NewDecoder(r).Decode(&reply); err != nil {
		return 0, err
	}

	total := 0.0
	for _, p := range reply.Payments {
		total += p.Amount
	}

	return total, nil

}

func main() {
	file, err := os.Open("reply.json")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	val, err := totalPayments(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println("total:", val)
}
