package main

import (
    "fmt"
    "time"

    "github.com/Firdous2307/go-task-timer/TaskTimer"
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
            taskTimer.StartTask(tasks)

        case 2:
            taskTimer.ViewTasks(tasks)
        case 3:
            fmt.Println("Goodbye! :)")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
