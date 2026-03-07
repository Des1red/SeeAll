package handlers

import (
	"SeeAll/internal/model"
	"encoding/json"
	"net/http"
)

type Config struct {
	API string `json:"api"`
}

func ConfigHandler(runtime model.Runtime) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var api string

		if runtime.Dev {
			api = runtime.APIDev + "/news"
		} else {
			api = runtime.APIProd + "/news"
		}

		cfg := Config{
			API: api,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cfg)
	}
}
