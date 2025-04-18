// api/queue/task.go
package queue

type Task struct {
	ID      int
	Payload string
}

var taskQueue = make(chan Task, 100)
var idCounter = 1

// AddTask adds a new task to the task queue with the specified payload.
// It creates a Task instance with a unique ID and the provided payload,
// increments the ID counter, and sends the task to the task queue.
//
// Parameters:
//   - payload: A string representing the data or information associated with the task.
func AddTask(payload string) {
	task := Task{ID: idCounter, Payload: payload}
	idCounter++
	taskQueue <- task
}

func GetQueue() <-chan Task {
	return taskQueue
}
