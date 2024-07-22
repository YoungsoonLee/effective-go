package main

import (
	"log"
	"net/http"
)

func v1HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK v1"))
}

func v1Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", v1HealthHandler)
	return mux
}

func v2HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK v2"))
}

func v2Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/_/health", v2HealthHandler)
	return mux
}

func main() {
	v1 := v1Mux()
	v2 := v2Mux()

	versionRouter := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-ver") == "2" {
			v2.ServeHTTP(w, r)
			return
		}

		v1.ServeHTTP(w, r)
	}

	http.HandleFunc("/", versionRouter)

	addr := ":8080"
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
