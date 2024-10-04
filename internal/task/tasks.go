package tasktimer

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func StartTask(tasks map[string]time.Duration) {
	var taskName string
	fmt.Print("Enter task name: ")

	// Use bufio to read the full line for task name
	reader := bufio.NewReader(os.Stdin)
	taskName, _ = reader.ReadString('\n')
	taskName = taskName[:len(taskName)-1]

	if taskName == "" {
		fmt.Println("Task name cannot be empty. Please try again.")
		return
	}

	fmt.Println("Task started. Press Enter to stop.")
	start := time.Now()
	_, err := fmt.Scanln() // Wait for Enter key
	if err != nil {
		// Handle error, eg. log it or return it
		fmt.Println("Error waiting for input:", err)
		return
	}
	duration := time.Since(start)

	tasks[taskName] = duration
	fmt.Printf("Task '%s' completed in %v\n", taskName, duration.Round(time.Second))
}

func ViewTasks(tasks map[string]time.Duration) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Completed Tasks:")
	for name, duration := range tasks {
		fmt.Printf("- %s: %v\n", name, duration.Round(time.Second))
	}
}
