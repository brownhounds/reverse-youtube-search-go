package api_error

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ApiError(w http.ResponseWriter, status int) {
	r := ErrorResponse{
		Message: http.StatusText(status),
		Status:  status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(r); err != nil {
		panic(err.Error())
	}
}
