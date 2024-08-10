//go:build prof

package main

import (
	_ "net/http/pprof"
)

func main() {}
