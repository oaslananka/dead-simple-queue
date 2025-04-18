// api/queue/worker.go
package queue

import (
	"dead-simple-queue/storage"
	"log"
)

// StartWorker initializes a background worker that continuously processes tasks
// from the queue. Each task is retrieved from the queue, logged, and then saved
// to storage. The worker runs in a separate goroutine to allow non-blocking
// execution.
func StartWorker() {
	go func() {
		for task := range GetQueue() {
			log.Printf("Processing task ID %d with payload %s", task.ID, task.Payload)
			storage.SaveTask(task.ID, task.Payload)
		}
	}()
}
