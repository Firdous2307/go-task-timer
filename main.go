package main

import (
    "fmt"
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
        fmt.Scan(&choice)

        switch choice {
        case 1:
            startTask(tasks)
        case 2:
            viewTasks(tasks)
        case 3:
            fmt.Println("Goodbye :) !")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func startTask(tasks map[string]time.Duration) {
    var taskName string
    fmt.Print("Enter task name: ")
    fmt.Scan(&taskName)

    