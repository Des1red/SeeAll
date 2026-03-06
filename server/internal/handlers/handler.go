package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"SeeAll/internal/sources"
)

func News(w http.ResponseWriter, r *http.Request) {

	if !validateGET(w, r) {
		return
	}

	// /news/greece → greece
	t := strings.TrimPrefix(r.URL.Path, "/news/")

	if t == "" {
		http.Error(w, "missing news type", http.StatusBadRequest)
		return
	}

	posts, err := sources.FetchByType(t)
	if err != nil {
		http.Error(w, "failed to fetch sources", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
