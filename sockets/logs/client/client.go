package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func sendFile(addr, fileName string) error {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("Failed to dial: %v", err)
	}
	defer c.Close()

	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Failed to get file info: %v", err)
	}

	// Send header line with file size & name
	_, err = fmt.Fprintf(c, "%s %d\n", fileName, fileInfo.Size())
	if err != nil {
		return fmt.Errorf("Failed to send file header: %v", err)
	}

	// Send file content
	_, err = io.Copy(c, file)
	if err != nil {
		return fmt.Errorf("Failed to send file: %v", err)
	}

	const maxReply = 1 << 10 // 1KB
	data, err := io.ReadAll(io.LimitReader(c, maxReply))
	if err != nil {
		return fmt.Errorf("Failed to read reply: %v", err)
	}

	reply := string(data)
	log.Printf("Reply: %s", reply)
	if strings.HasPrefix(reply, "error") {
		return fmt.Errorf("Server error: %s", reply)
	}

	return nil
}

func main() {
	if err := sendFile("localhost:8765", "client.go"); err != nil {
		log.Panicf("Failed to send file: %v", err)
	}
}
