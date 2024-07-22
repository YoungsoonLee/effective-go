package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	// mime type -> encoder function
	registry = make(map[string]Encoder)
)

// Encoder function
type Encoder func(w io.Writer, metrics []Metric) error

// Register registers an encoder function for a given mime type
// It'll panic if the mime type is already registered
func Register(mimeType string, enc Encoder) {
	if _, ok := registry[mimeType]; ok {
		panic("encoder already registered")
	}
	registry[mimeType] = enc
}

// EncodeJSON encodes the metrics as JSON
func EncodeJSON(w io.Writer, metrics []Metric) error {
	return json.NewEncoder(w).Encode(metrics)
}

// EncodeCSV encodes the metrics as CSV
func EncodeCSV(w io.Writer, metrics []Metric) error {
	wtr := csv.NewWriter(w)
	// Write the header
	if err := wtr.Write([]string{"time", "name", "value"}); err != nil {
		return err
	}

	// Record to write
	r := make([]string, 3)
	for _, m := range metrics {
		r[0] = m.Time.Format(time.RFC3339)
		r[1] = m.Name
		r[2] = fmt.Sprintf("%f", m.Value)
		if err := wtr.Write(r); err != nil {
			return err
		}
	}

	wtr.Flush()

	return nil
}

// init() registers the JSON and CSV encoders
func init() {
	Register(jsonMimeType, EncodeJSON)
	Register(csvMimeType, EncodeCSV)
}

const (
	jsonMimeType = "application/json"
	csvMimeType  = "text/csv"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "missing query", http.StatusBadRequest)
		return
	}
	mimeType := requestMimeType(r)
	enc, ok := registry[mimeType]
	if !ok {
		http.Error(w, "unsupported mime type", http.StatusBadRequest)
		return
	}

	log.Printf("requesting %q with %q", query, mimeType)

	metrics, err := queryDB(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", mimeType)
	if err := enc(w, metrics); err != nil {
		size := len(metrics)
		const format = "can't encode %d metrics with %q - %s"
		log.Printf(format, size, mimeType, err)
	}

}

func requestMimeType(r *http.Request) string {
	// Get mime type from HTTP Accept header
	mimeType := r.Header.Get("Accept")
	if mimeType == "" || mimeType == "*/*" {
		return jsonMimeType // default to JSON
	}
	return mimeType
}

// Metric represents a single data point
type Metric struct {
	Time  time.Time
	Name  string
	Value float64
}

func queryDB(query string) ([]Metric, error) {
	log.Printf("querying DB with %q", query)
	return nil, nil
}

func main() {
	http.HandleFunc("/metrics", queryHandler)
	addr := ":8090"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
