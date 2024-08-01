package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func newID() string {
	return uuid.New().String()
}

type ctxKey string // !!!

const (
	valuesKey ctxKey = "ctxArgs"
)

// Values is a struct to hold request ID and logger
type Values struct {
	RequestID string
	Logger    *log.Logger
}

func idLogger(id string) *log.Logger {
	prefix := fmt.Sprintf("[%s] %s", id, log.Prefix())
	return log.New(log.Writer(), prefix, log.Flags())
}

func ctxLogger(ctx context.Context) *log.Logger {
	v, ok := ctx.Value(valuesKey).(*Values)
	if !ok {
		return stdLogger()
	}

	if v.Logger == nil {
		panic(fmt.Sprintf("nil logger in %#v", v))
	}

	return v.Logger
}

// stdLogger returns a logger behaves like the to-level function in
// the log package
func stdLogger() *log.Logger {
	return log.New(log.Writer(), log.Prefix(), log.Flags())
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	timeout := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	id := newID()
	logger := idLogger(id)
	logger.Printf("info: usersHandler: id=%s", id)

	values := Values{
		RequestID: id,
		Logger:    logger,
	}
	ctx = context.WithValue(ctx, valuesKey, &values)
	users, err := getAllUsers(ctx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logger.Printf("info: usersHandler: got %d users", len(users))

	json.NewEncoder(w).Encode(users)
}

func getAllUsers(ctx context.Context) ([]string, error) {
	logger := ctxLogger(ctx)
	logger.Printf("info: getting all users")

	// FIXME: Connect to database, query ...
	return nil, nil
}

func main() {
	http.HandleFunc("/users", usersHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}
