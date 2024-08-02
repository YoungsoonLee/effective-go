package config

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	cfgOnce sync.Once
	// Config is the configuration for the application
	Config struct {
		ListenAddr string
		Verbose    bool
	}
)

func loadConfig(envPrefix string) {
	// Utility function to get from environment with prefix
	getEnv := func(name string) string {
		key := fmt.Sprintf("%s_%s", envPrefix, name)
		return os.Getenv(key)
	}

	addr := getEnv("ADDRESS")
	if len(addr) == 0 {
		Config.ListenAddr = ":8080" // default
	} else {
		Config.ListenAddr = addr
	}

	verbose := getEnv("VERBOSE")
	if verbose == "1" || verbose == "yes" || verbose == "on" {
		Config.Verbose = true
	}

	log.Printf("configuration loaded (prefix: %s): %+v", envPrefix, Config)
}

// LoadConfig loads the configuration once from environment variables
func LoadConfig(envPrefix string) {
	cfgOnce.Do(func() {
		loadConfig(envPrefix)
	})
}
