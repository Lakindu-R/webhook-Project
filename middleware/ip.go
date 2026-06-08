package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {

	ip := r.Header.Get("X-Forwarded-For")

	if ip == "" {
		ip = r.RemoteAddr
		fmt.Println("RemorteAddr",ip)
	}

	// remove port (127.0.0.1:54321 → 127.0.0.1)
	if strings.Contains(ip, ":") {
		host, _, err := net.SplitHostPort(ip)
		if err == nil {
			return host
		}
	}

	return ip
}