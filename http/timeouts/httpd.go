package main

import (
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
