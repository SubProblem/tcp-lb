package balancer

import (
	"encoding/json"
	"net/http"
)

type StatusResponse struct {
	Address           string `json:"address"`
	Healthy           bool   `json:"healthy"`
	ActiveConnections int32  `json:"active_connections"`
	TotalRequests     uint64 `json:"total_requests"`
}

func NewStatusResponseHandler(backends []Backend) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusResponses := make([]StatusResponse, 0)
		for i := range backends {
			response := StatusResponse{
				Address:           backends[i].Address,
				Healthy:           backends[i].Healthy.Load(),
				ActiveConnections: backends[i].ActiveConns.Load(),
				TotalRequests:     backends[i].TotalRequests.Load(),
			}
			statusResponses = append(statusResponses, response)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statusResponses); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}