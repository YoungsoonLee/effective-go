package main

import (
	"net/http"
	"strings"
)

//go:generate go run _scripts/gen_ips.go _scripts/allowed_ips.txt ips.go
//go:generate go fmt ips.go

func requestIP(r *http.Request) string {
	field := strings.Split(r.RemoteAddr, ":")
	if len(field) != 2 {
		return ""
	}
	return field[0]
}

func main() {
	if ip := requestIP(r); !IPAllowed(ip) {
		http.Error(w, "not allowed", http.StatusForbidden)
		return
	}
}
