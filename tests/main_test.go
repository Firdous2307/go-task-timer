package tests

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"

	tasktimer "github.com/Firdous2307/go-task-timer/internal/task"
)

func TestStartTask(t *testing.T) {
	tasks := make(map[string]time.Duration)

	// Simulate input for task name
	input := "Test Task\n"
	reader := strings.NewReader(input)

	// Create a temporary file to simulate stdin
	tempFile, err := os.CreateTemp("", "testinput")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write input to the temp file and then read from it
	_, err = io.Copy(tempFile, reader)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Reset the file pointer to the beginning
	tempFile.Seek(0, 0)

	// Redirect stdin to the temp file
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = tempFile

	// Start the task
	go tasktimer.StartTask(tasks)

	// Allow some time for the task to run
	time.Sleep(2 * time.Second)
	os.Stdin.Close() // Simulate pressing Enter

	// Check if task duration is greater than 0
	if duration, exists := tasks["Test Task"]; !exists || duration < 0 {
		t.Errorf("Expected task 'Test Task' duration to be greater than 0 seconds, got %v", duration)
	}
}

func TestViewTasks(t *testing.T) {
	tasks := map[string]time.Duration{
		"Task 1": 2 * time.Second,
		"Task 2": 3 * time.Second,
	}

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// View the tasks
	tasktimer.ViewTasks(tasks)
	w.Close()

	// Read output from pipe
	var buf strings.Builder
	io.Copy(&buf, r)
	os.Stdout = oldStdout

	// Check if the output contains the expected task durations
	output := buf.String()
	if !strings.Contains(output, "Task 1: 2s") || !strings.Contains(output, "Task 2: 3s") {
		t.Errorf("Expected output to contain task durations, got: %s", output)
	}
}
