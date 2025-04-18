package main

import (
	"dead-simple-queue/api"
	"dead-simple-queue/queue"
	"log"
	"net/http"
)

// main is the entry point of the application. It starts the worker process
// for handling queued tasks and sets up an HTTP server to listen for incoming
// requests. The server provides an endpoint at "/enqueue" to handle task
// enqueuing. The application listens on port 8080.
func main() {
	queue.StartWorker()

	http.HandleFunc("/enqueue", api.EnqueueHandler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
