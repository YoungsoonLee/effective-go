package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Message is a channel message.
type Message struct {
	From    string `json:"from"`
	Channel string `json:"channel"`
	Body    string `json:"body"`
	ID      string `json:"-"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	m.ID = uuid.NewString()
	// Adding code redacted

	resp := map[string]any{
		"id":   m.ID,
		"time": time.Now().UTC(),
	}
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

func main() {

}
