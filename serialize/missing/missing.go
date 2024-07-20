package main

import (
	"encoding/json"
	"errors"
	"time"
)

// Log is a log event.
type Log struct {
	Time    time.Time
	Level   int
	Message string
}

func parseLog(data []byte) (*Log, error) {
	var l struct {
		Time    *time.Time // pointer to time.Time !!!
		Level   *int       // pointer to int !!!
		Message string
	}

	if err := json.Unmarshal(data, &l); err != nil {
		return nil, err
	}

	if l.Time == nil {
		return nil, errors.New("missing Time")
	}

	if l.Level == nil {
		return nil, errors.New("missing Level")
	}

	return &Log{
		Time:    *l.Time,
		Level:   *l.Level,
		Message: l.Message,
	}, nil
}
