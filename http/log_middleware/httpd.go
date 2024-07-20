package main

import (
	"log"
	"net/http"
	"time"
)

// addLoging is a middleware that logs the request to the standard output.
func addLogging(name string, handler http.Handler) http.Handler {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("%s took %s", name, duration)
	}

	return http.HandlerFunc(wrapper)
}

// queryHandler is a handler that simulates a query.
func queryHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("query results"))
}

func main() {
	hdlr := addLogging("query", http.HandlerFunc(queryHandler))
	http.Handle("/query", hdlr)

}
