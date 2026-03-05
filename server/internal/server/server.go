package server

import (
	"log"
	"net/http"

	"SeeAll/internal/handlers"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if origin == "https://des1red.github.io" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Vary", "Origin")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Start(addr string) {

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/news/daily", handlers.Daily)
	mux.HandleFunc("/news/live", handlers.Live)

	log.Println("server running on", addr)

	err := http.ListenAndServe(addr, enableCORS(mux))
	if err != nil {
		log.Fatal(err)
	}
}
