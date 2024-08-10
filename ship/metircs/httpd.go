package main

import (
	"expvar"
	"fmt"
	"net/http"
	"strconv"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func addMetrics(name string, h http.Handler) http.Handler {
	calls := expvar.NewInt(fmt.Sprintf("%s.calls", name))
	errors := expvar.NewInt(fmt.Sprintf("%s.errors", name))
	oks := expvar.NewInt(fmt.Sprintf("%s.oks", name))

	fn := func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)

		sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		h.ServeHTTP(sw, r)

		if sw.statusCode >= http.StatusBadRequest {
			errors.Add(1)
		} else {
			oks.Add(1)
		}
	}

	return http.HandlerFunc(fn)
}

func main() {
	h := addMetrics("lookup", http.HandlerFunc(lookupHandler))
	http.Handle("/lookup", h)
}

func lookupHandler(w http.ResponseWriter, r *http.Request) {
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil || lat < -90 || lat > 90 {
		http.Error(w, "invalid lat", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil || lon < -180 || lon > 180 {
		http.Error(w, "invalid lon", http.StatusBadRequest)
		return
	}

	// Lookup code redacted
	fmt.Fprintf(w, "lat=%f, lon=%f", lat, lon)
}
