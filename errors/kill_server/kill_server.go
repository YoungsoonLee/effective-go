package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return fmt.Errorf("failed to open pid file: %w", err)
	}
	defer file.Close()

	var pid int
	_, err = fmt.Fscanf(file, "%d", &pid)
	if err != nil {
		return fmt.Errorf("bad PID in %q: %w", pidFile, err)
	}

	if err := os.Remove(pidFile); err != nil {
		log.Printf("can't remove %q - %s", pidFile, err) // warn, no error
	}

	return kill(pid)
}

func kill(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return proc.Kill()
}

func killFromFiles(logFiles []string) error {
	for _, file := range logFiles {
		err := killServer(file)
		if err == nil {
			return nil
		}

		if !errors.Is(err, os.ErrNotExist) { // file not found
			return err
		}
	}

	files := strings.Join(logFiles, ", ")
	return fmt.Errorf("no server PIDs found in %s", files)
}
