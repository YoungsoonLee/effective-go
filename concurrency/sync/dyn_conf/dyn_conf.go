package dynconf

import (
	"sync"
	"time"
)

/*
sync.RWMutex will allow only one writer with no readers or many readers with no writer.
You can use a regular sync.Mutex, but then only one reader at a time will be able to access the configuration.
This is not an optimal use of the concurrency Go offers.
*/

var (
	cfgLock sync.RWMutex
	config  = make(map[string]string)
)

// GetConfig returns the value for key in configuration.
func GetConfig(key string) string {
	cfgLock.RLock()
	defer cfgLock.RUnlock()
	return config[key]
}

// ReloadConfig reloads configuration.
func ReloadConfig() {
	cfgLock.Lock()
	defer cfgLock.Unlock()
	config["updated"] = time.Now().String()
	// TODO: finish loading configuration
}
