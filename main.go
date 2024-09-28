package main

import (
    "fmt"
    "html/template"
    "sync"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Firdous2307/go-task-timer/TaskTimer"
)

type Task struct {
    Name     string
    Duration time.Duration
}

var (
    tasks []Task
    mu    sync.Mutex
)

func main() {
    app := fiber.New()

    // Middleware for logging requests
    app.Use(logger.New())

    // Serve static files
    app.Static("/static", "./static")


    // Web server routes
    app.Get("/", homeHandler)
    app.Post("/start", startTaskHandler)
    app.Post("/stop", stopTaskHandler)

    // Run web server in a goroutine to allow CLI usage alongside
    go func() {
        app.Listen(":8080")
    }()

    // CLI loop for user input
    cliLoop()
}


func homeHandler(c *fiber.Ctx) error {
    mu.Lock()
    defer mu.Unlock()

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    return tmpl.Execute(c, struct{ Tasks []Task }{Tasks: tasks})
}

func startTaskHandler(c *fiber.Ctx) error {
    taskName := c.FormValue("task")
    if taskName == "" {
        return c.Status(400).SendString("Task name is required.")
    }

    mu.Lock()
    tasks = append(tasks, Task{Name: taskName, Duration: 0}) // Start task without duration
    mu.Unlock()

    return c.SendString("Task started")
}

func stopTaskHandler(c *fiber.Ctx) error {
    taskName := c.FormValue("task")
    durationStr := c.FormValue("duration")
    if taskName == "" || durationStr == "" {
        return c.Status(400).SendString("Task name and duration are required.")
    }

    duration, err := time.ParseDuration(durationStr + "s") // Convert duration to time.Duration
    if err != nil {
        return c.Status(400).SendString("Invalid duration.")
    }

    mu.Lock()
    for i, task := range tasks {
        if task.Name == taskName && task.Duration == 0 {
            tasks[i].Duration = duration // Update the duration for the task
            break
        }
    }
    mu.Unlock()

    return c.SendString("Task stopped")
}

func cliLoop() {
    cliTasks := make(map[string]time.Duration)

    for {
        fmt.Println("\n1. Start a task")
        fmt.Println("2. View tasks")
        fmt.Println("3. Exit")
        fmt.Print("Choose an option: ")

        var choice int
        _, err := fmt.Scan(&choice)
        if err != nil {
            fmt.Println("Invalid choice. Please try again.")
            var discard string
            fmt.Scanln(&discard)
            continue
        }

        switch choice {
        case 1:
            TaskTimer.StartTask(cliTasks)
        case 2:
            TaskTimer.ViewTasks(cliTasks)
        case 3:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
