package main

// IPAllowed returns true if connections from ip are allowed
func IPAllowed(ip string) bool {
	return allowedIPs[ip]
}

var allowedIPs = map[string]bool{
	"127.0.0.1":    true,
	"79.180.16.99": true,
}
