package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Level byte

const (
	Debug Level = iota + 1
	Info
	Error
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Error:
		return "ERROR"
	}

	return fmt.Sprintf("Level <%d>", l)
}

type syncer interface {
	Sync() error
}

// Logger is a logger.
type Logger struct {
	level Level
	w     io.Writer
	s     syncer
}

type nopSyncer struct{}

func (ns nopSyncer) Sync() error {
	return nil
}

// END NopSyncer

// NewLogger returns a new logger with level and output.
func NewLogger(level Level, out io.Writer) Logger {
	log := Logger{level, out, nopSyncer{}}
	if s, ok := out.(syncer); ok {
		log.s = s
	}
	return log
}

func (l Logger) log(level Level, format string, args ...any) {
	if l.level > level {
		return
	}

	msg := fmt.Sprintf(format, args...)
	ts := time.Now().UTC().Format(time.RFC3339)
	fmt.Fprintf(l.w, "[%s] - %s - %s\n", ts, level, msg)

	l.s.Sync()
}

// Debug logs a debug message.
func (l Logger) Debug(format string, args ...any) {
	l.log(Debug, format, args...)
}

// Info logs an info message.
func (l Logger) Info(format string, args ...any) {
	l.log(Info, format, args...)
}

// Error logs an error message.
func (l Logger) Error(format string, args ...any) {
	l.log(Error, format, args...)
}

func main() {
	log := NewLogger(Info, os.Stdout)
	log.Debug("debug")
	log.Info("info")
	log.Error("error")
}
