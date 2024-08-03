package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strings"
)

var maxSize int64 = 1024 * 1024

func main() {
	addr := ":8765"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		log.Printf("Accepted connection from %s", conn.RemoteAddr())

		go fileHandler(conn)
	}
}

func fileHandler(conn net.Conn) {
	defer conn.Close()

	var fileName string
	var size int64
	if _, err := fmt.Fscanf(conn, "%s %d", &fileName, &size); err != nil {
		log.Printf("Failed to read file header: %v", err)
		fmt.Fprintf(conn, "Failed to read file header: %v\n", err)
		return
	}
	log.Printf("Receiving file %s (%d bytes)", fileName, size)

	if size > maxSize {
		log.Printf("File is too large: %d bytes. maxSize: %d", size, maxSize)
		fmt.Fprintf(conn, "File is too large: %d bytes\n", size)
		return
	}

	// Save in "logs" directory
	fileName = path.Join("logs", strings.TrimSpace(fileName))
	log.Printf("Saving to %s", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		fmt.Fprintf(conn, "Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	n, err := io.CopyN(file, conn, size)
	if err != nil {
		log.Printf("Failed to save file: %v", err)
		fmt.Fprintf(conn, "Failed to save file: %v\n", err)
		return
	}

	if n != size {
		log.Printf("Failed to save file: expected %d bytes, saved %d", size, n)
		fmt.Fprintf(conn, "Failed to save file: expected %d bytes, saved %d\n", size, n)
		return
	}

	log.Printf("File saved successfully")
}
