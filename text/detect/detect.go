package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"

	"golang.org/x/net/html/charset"
)

// ctypeEncoding gets the encoding from HTTP Content-Type header.
func ctypeEncoding(ctype string) string {
	_, params, err := mime.ParseMediaType(ctype)
	if err != nil {
		return ""
	}

	return params["charset"]
}

func dataEncoding(data []byte) string {
	_, name, certain := charset.DetermineEncoding(data, "text/plain")
	if certain {
		return name
	}

	return ""
}

func main() {
	resp, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	enc := ctypeEncoding(resp.Header.Get("Content-Type"))
	log.Printf("content-type encoding: %s", enc)
	if enc != "" {
		fmt.Printf("content-type encoding: %s\n", enc)
		os.Exit(0)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	enc = dataEncoding(data)
	if enc != "" {
		fmt.Printf("data encoding: %s\n", enc)
		os.Exit(0)
	}

	fmt.Println("unknown encoding")
	os.Exit(1)

}
