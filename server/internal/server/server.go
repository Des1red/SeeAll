package server

import (
	"log"
	"net/http"
	"strings"

	"SeeAll/internal/database/store"
	"SeeAll/internal/handlers"
	"SeeAll/internal/model"

	"github.com/Des1red/goauthlib/goauth"
)

func enableCORS(runtime model.Runtime, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Admin routes do not use CORS
		if strings.HasPrefix(r.URL.Path, "/admin") {
			next.ServeHTTP(w, r)
			return
		}
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
	goauth.JWTSecret([]byte(runtime.JWTsecret))
	// use db
	goauth.UseStore(store.NewAuthTokenStore())

	// Optional cookie config
	goauth.Cookies(goauth.CookieConfig{
		Secure:   !runtime.Dev,
		SameSite: http.SameSiteStrictMode,
	})
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/news/", handlers.News)
	mux.HandleFunc("/config", handlers.ConfigHandler(runtime))

	// admin routes
	mux.HandleFunc("/admin/login", handlers.AdminLogin(runtime))
	mux.HandleFunc("/admin/logout", handlers.AdminLogout)
	mux.Handle("/admin/api/stats", goauth.Admin(handlers.AdminStats))
	mux.Handle("/admin", goauth.Admin(handlers.AdminPage))

	log.Println("server running on", runtime.Port)
	err := http.ListenAndServe(
		runtime.Port,
		enableCORS(runtime, metricsMiddleware(runtime, mux)),
	)
	if err != nil {
		log.Fatal(err)
	}
}
