package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// Location is location on Earth
type Location struct {
	Lat, Lng float64
}

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	m, err := unix.Mmap(
		int(file.Fd()), 0, int(fi.Size()),
		unix.PROT_READ, unix.MAP_PRIVATE,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer unix.Munmap(m)

	pos := 0
	locPrefix := []byte("loc:{")
	var loc Location
	for {
		i := bytes.Index(m[pos:], locPrefix) // find "loc:{"
		if i == -1 {
			break
		}

		i += len(locPrefix) - 1 //move over "loc:{"
		start := pos + i
		size := bytes.IndexByte(m[start:], '}')
		if size == -1 {
			break
		}

		size++
		if err := json.Unmarshal(m[start:start+size], &loc); err != nil {
			log.Println(err)
		}

		fmt.Printf("Location: %v\n", loc)
		pos = start + size + 1 // move after end of current JSON document
	}
}
