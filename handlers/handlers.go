package handlers

import (
	"encoding/json"
	"go-reverse-youtube-search/api_error"
	"go-reverse-youtube-search/clients/youtube"
	"net/http"
	"strings"
)

func YoutubeSearchHandler(w http.ResponseWriter, r *http.Request) {
	searchParam := r.URL.Query().Get("q")

	if strings.TrimSpace(searchParam) == "" {
		api_error.ApiError(w, http.StatusBadRequest)
		return
	}

	err, res := youtube.Search(searchParam)

	if err != nil {
		api_error.ApiError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		api_error.ApiError(w, http.StatusInternalServerError)
	}
}
