package server

import (
	"crypto/sha256"
	"encoding/hex"
	"net"
	"net/http"
	"strings"

	"SeeAll/internal/metrics"
	"SeeAll/internal/model"
)

func metricsMiddleware(runtime model.Runtime, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/admin") || strings.HasPrefix(r.URL.Path, "/config") || strings.HasPrefix(r.URL.Path, "/ping") {
			next.ServeHTTP(w, r)
			return
		}
		// Don't count preflight, favicon, or static assets
		if r.Method == http.MethodOptions ||
			r.URL.Path == "/favicon.ico" ||
			r.URL.Path == "/" ||
			strings.HasSuffix(r.URL.Path, ".html") ||
			strings.HasSuffix(r.URL.Path, ".css") ||
			strings.HasSuffix(r.URL.Path, ".js") ||
			strings.HasSuffix(r.URL.Path, ".png") ||
			strings.HasSuffix(r.URL.Path, ".json") {
			next.ServeHTTP(w, r)
			return
		}

		ip := extractIP(r)
		hash := hashIP(ip, runtime.HashSalt)

		metrics.IncVisitor(hash)
		metrics.IncActive()
		metrics.IncTotal()
		metrics.IncEndpoint(r.URL.Path)

		defer metrics.DecActive()

		next.ServeHTTP(w, r)
	})
}

func extractIP(r *http.Request) string {

	// Reverse proxy headers first
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	// fallback
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	return host
}

func hashIP(ip string, hash string) string {
	h := sha256.Sum256([]byte(ip + hash))
	return hex.EncodeToString(h[:])
}
