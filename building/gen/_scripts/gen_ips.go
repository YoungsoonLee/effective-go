package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var header = ` package main
// IPAllowed returns true if connections from ip are allowed
func IPAllowed(ip string) bool {
    return allowedIPs[ip]
}
var allowedIPs = map[string]bool{`

func main() {
	in, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Fprintln(out, header)

	s := bufio.NewScanner(in)
	for s.Scan() {
		fmt.Fprintf(out, "%q: true,\n", s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(out, "}")
}
