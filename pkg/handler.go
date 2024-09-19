package pkg

import (
	"context"

	"encoding/json"
	"net/http"
	"time"
)

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	select {
	case <-ctx.Done():
		http.Error(w, "Request Timeout", http.StatusRequestTimeout)
		return
	default:
		res := EchoResponse{
			Message:   req.Message,
			Timestamp: time.Now(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
