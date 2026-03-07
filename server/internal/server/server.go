package server

import (
	"log"
	"net/http"

	"SeeAll/internal/handlers"
	"SeeAll/internal/model"
)

func enableCORS(runtime model.Runtime, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if runtime.Dev && origin == runtime.DevOrigin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		if origin == runtime.ProdOrigin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Vary", "Origin")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Start(runtime model.Runtime) {

	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/news/", handlers.News)
	mux.HandleFunc("/config", handlers.ConfigHandler(runtime))
	log.Println("server running on", runtime.Port)

	err := http.ListenAndServe(runtime.Port, enableCORS(runtime, mux))
	if err != nil {
		log.Fatal(err)
	}
}
