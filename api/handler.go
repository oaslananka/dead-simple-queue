// api/handler.go
package api

import (
	"dead-simple-queue/queue"
	"encoding/json"
	"net/http"
)

// EnqueueRequest represents the structure of a request to enqueue a new item
// into the queue. It contains a single field, Payload, which holds the data
// to be added to the queue in JSON format.
type EnqueueRequest struct {
	Payload string `json:"payload"`
}

// EnqueueHandler handles HTTP POST requests to enqueue a task into the queue.
// It expects a JSON payload in the request body that conforms to the EnqueueRequest structure.
// If the request method is not POST, it responds with a 405 Method Not Allowed status.
// If the request body cannot be decoded into the EnqueueRequest structure, it responds with a 400 Bad Request status.
// On successful decoding, it adds the task to the queue and responds with a 202 Accepted status.
func EnqueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EnqueueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	queue.AddTask(req.Payload)
	w.WriteHeader(http.StatusAccepted)
}
