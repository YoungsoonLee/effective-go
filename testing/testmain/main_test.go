package main

import (
	"log"
	"os"
	"testing"

	"gopkg.in/yaml.v2"
)

const (
	envConfigKey = "APP_CONFIG_FILE"
)

func setupTest() error {
	file, err := os.CreateTemp("", "*.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	cfg := map[string]any{
		"verbose": true,
		"dsn":     "postgres://localhost:5432",
	}

	if err := yaml.NewEncoder(file).Encode(cfg); err != nil {
		return err
	}

	os.Setenv(envConfigKey, file.Name())
	log.Printf("Config file: %q", file.Name())
	return nil
}

func teardownTest() {
	fileName := os.Getenv(envConfigKey)
	if err := os.Remove(fileName); err != nil {
		log.Printf("Failed to remove config file %q: %v", fileName, err)
	}
}

func runTests(m *testing.M) int {
	if err := setupTest(); err != nil {
		log.Printf("Failed to setup test: %v", err)
		return 1
	}
	defer teardownTest()

	return m.Run()
}

func TestMain(m *testing.M) {
	code := runTests(m)
	os.Exit(code)
}
