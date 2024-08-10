package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ardanlabs/conf/v3"
)

var cfg struct {
	Web struct {
		Addr string `conf:"default::8080,env:ADDR"`
	}
}

func main() {
	help, err := conf.Parse("APP", &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			os.Exit(0)
		}
		log.Fatalf("error: parsing config: %v", err)
	}

	if err := validateAddr(cfg.Web.Addr); err != nil {
		log.Fatalf("error: invalid web addr: %v", err)
	}
}

func validateAddr(addr string) error {
	i := strings.Index(addr, ":")
	if i == -1 {
		return fmt.Errorf("%q: missing port in address", addr)
	}
	port, err := strconv.Atoi(addr[i+1:])
	if err != nil {
		return fmt.Errorf("%q: invalid port in address: %v", addr, err)
	}

	const maxPort = 65_535
	if port < 0 || port > maxPort {
		return fmt.Errorf("%q: port outside valid range (0-%d)", addr, maxPort)
	}

	return nil
}
