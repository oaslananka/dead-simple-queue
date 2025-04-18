// storage/file.go
package storage

import (
	"fmt"
	"os"
)

// SaveTask saves a task with the given ID and payload to a log file named "tasks.log".
// If the file does not exist, it will be created. If it exists, the task will be appended
// to the file. The function handles errors that may occur during file operations.
//
// Parameters:
//   - id: An integer representing the unique identifier of the task.
//   - payload: A string containing the task's payload or description.
//
// Note: The function writes the task information in the format:
//
//	"Task ID: <id>, Payload: <payload>"
func SaveTask(id int, payload string) {
	f, err := os.OpenFile("tasks.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "Task ID: %d, Payload: %s\n", id, payload)
}
