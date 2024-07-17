package main

import (
	"bufio"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func numRedirects(r io.Reader) (nLines int, nRedirects int, err error) {
	s := bufio.NewScanner(r)
	//nLines, nRedirects := 0, 0
	for s.Scan() {
		nLines++
		// Example:
		// 203.252.212.44 -- [2020-03-01T14:32:05Z] "GET / HTTP/1.1" 200 107 \
		// "GET /ksc.html HTTP/1.1" 200 167
		fields := strings.Fields(s.Text())
		code := fields[len(fields)-2] // code is one before the last field
		if code[0] == '3' {           // 3xx is a redirect
			nRedirects++
		}
	}

	if err := s.Err(); err != nil {
		return -1, -1, err
	}

	return nLines, nRedirects, nil
}

func main() {
	matches, err := filepath.Glob("logs/http-*.log")
	if err != nil {
		log.Fatal(err)
	}

	nLines, nRedirects := 0, 0
	for _, name := range matches {
		file, err := os.Open(name)
		if err != nil {
			log.Print(err)
			continue
		}

		var r io.Reader = file
		if strings.HasSuffix(name, ".gz") {
			r, err = gzip.NewReader(r)
			if err != nil {
				log.Print(err)
				continue
			}
		}

		lines, redirects, err := numRedirects(r)
		if err != nil {
			log.Print(err)
		}

		nLines += lines
		nRedirects += redirects
	}

	log.Printf("lines: %d, redirects: %d", nLines, nRedirects)
}
