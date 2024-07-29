package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type errWriter struct {
	http.ResponseWriter
	statusCode int
}

func (ew *errWriter) WriteHeader(statusCode int) {
	ew.statusCode = statusCode
	ew.ResponseWriter.WriteHeader(statusCode)
}

func logErr(w http.ResponseWriter, r *http.Request) {
	ew := errWriter{ResponseWriter: w}
	http.DefaultServeMux.ServeHTTP(&ew, r)

	if ew.statusCode >= http.StatusBadRequest {
		// log error
		log.Printf("error: %s %s %d", r.Method, r.URL.Path, ew.statusCode)
	}

}

func tzOffset(name string) (int, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return 0, err
	}

	_, offset := time.Now().In(loc).Zone()
	return offset, nil
}

const (
	offsetPrefix = "/offset/"
)

func offsetHandler(w http.ResponseWriter, r *http.Request) {
	// "/offset/US/Pacific" -> "US/Pacific"
	name := r.URL.Path[len(offsetPrefix):]
	offset, err := tzOffset(name)
	if err != nil {
		errMsg := fmt.Sprintf("unknown offset - %q", name)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "%d\n", offset)
}

func main() {
	http.HandleFunc(offsetPrefix, offsetHandler)

	handler := http.HandlerFunc(logErr)
	if err := http.ListenAndServe(":8088", handler); err != nil {
		log.Fatalf("could not listen on 8080: %v", err)
	}
}
