// tests/queue_test.go
package tests

import (
	"dead-simple-queue/queue"
	"testing"
)

// TestAddTask tests the AddTask function of the queue.
// It verifies that a task added to the queue has the expected payload
// by retrieving the task from the queue and comparing its payload
// with the input payload. If the payloads do not match, the test fails.
func TestAddTask(t *testing.T) {
	taskPayload := "unit-test-payload"
	queue.AddTask(taskPayload)
	task := <-queue.GetQueue()

	if task.Payload != taskPayload {
		t.Errorf("Expected payload %s but got %s", taskPayload, task.Payload)
	}
}
