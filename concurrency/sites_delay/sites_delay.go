package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Info is information on site
type Info struct {
	url        string
	statusCode int
	delay      time.Duration
	err        error
}

func siteInfo(url string) (Info, error) {
	var info Info

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()

	// Consume data from server
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		return info, err
	}

	info.delay = time.Since(start)
	info.statusCode = resp.StatusCode

	return info, nil
}

func sitesInfo(urls []string) (map[string]Info, error) {
	// out := make(map[string]Info)
	// for _, url := range urls {
	// 	info, err := siteInfo(url)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	out[url] = info
	// }

	// return out, nil

	ch := make(chan Info)
	for _, url := range urls {
		go func(url string) {
			info, err := siteInfo(url)
			info.err = err
			info.url = url
			ch <- info
		}(url)
	}

	out := make(map[string]Info)
	for range urls {
		info := <-ch
		if info.err != nil {
			return nil, info.err
		}
		out[info.url] = info
	}

	return out, nil
}

func main() {
	start := time.Now()

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
	}

	infos, err := sitesInfo(urls)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for url, info := range infos {
		fmt.Printf("%s: %+v\n", url, info)
	}

	duration := time.Since(start)
	fmt.Printf("Total time: %v\n", duration)

}
