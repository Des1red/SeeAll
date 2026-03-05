package handlers

import (
	"encoding/json"
	"net/http"

	"SeeAll/internal/sources"
)

func Live(w http.ResponseWriter, r *http.Request) {
	if !validateGET(w, r) {
		return
	}
	posts, err := sources.FetchByType("live")
	if err != nil {
		http.Error(w, "failed to fetch sources", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
