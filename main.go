package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    tasks := make(map[string]time.Duration)

    for {
        fmt.Println("\n1. Start a task")
        fmt.Println("2. View tasks")
        fmt.Println("3. Exit")
        fmt.Print("Choose an option: ")

        var choice int
        _, err := fmt.Scan(&choice)
        if err != nil {
            fmt.Println("Invalid choice. Please try again.")
            // Clear the input buffer if there's an error
            var discard string
            fmt.Scanln(&discard)
            continue
        }

        switch choice {
        case 1:
            startTask(tasks)

        // case 2:
        //     viewTasks(tasks)
        // case 3:
        //     fmt.Println("Goodbye! :)")
        //     return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func startTask(tasks map[string]time.Duration) {
    var taskName string
    fmt.Print("Enter task name: ")

    // Use bufio to read the full line for task name
    reader := bufio.NewReader(os.Stdin)
    taskName, _ = reader.ReadString('\n')
    taskName = taskName[:len(taskName)-1] // Remove the newline character

    if taskName == "" {
        fmt.Println("Task name cannot be empty. Please try again.")
        return
    }

    fmt.Println("Task started. Press Enter to stop.")
    start := time.Now()
    fmt.Scanln() // Wait for Enter key
    duration := time.Since(start)

    tasks[taskName] = duration
    fmt.Printf("Task '%s' completed in %v\n", taskName, duration.Round(time.Second))
}
