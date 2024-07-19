package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// Location is location on Earth.
type Location struct {
	Lat float64
	Lng float64
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	m, err := unix.Mmap(
		int(file.Fd()), 0, int(fi.Size()),
		unix.PROT_READ, unix.MAP_PRIVATE,
	)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer unix.Munmap(m) // free the mmap when done with it

	pos := 0 // current position in data
	locPrefix := []byte("loc:{")
	var loc Location
	for {
		i := bytes.Index(m[pos:], locPrefix) // find "loc:"
		if i == -1 {
			break
		}

		i += len(locPrefix) - 1 // move over "loc:"
		start := pos + i
		size := bytes.IndexByte(m[start:], '}') // find closing }
		if size == -1 {
			break
		}
		size++ // move over }

		if err := json.Unmarshal(m[start:start+size], &loc); err != nil {
			log.Fatalf("error: %s", err)
		}
		fmt.Printf("%+v\n", loc)
		pos = start + size + 1 // move after end of current JSON document
	}
}
