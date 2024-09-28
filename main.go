package main

import (
    "fmt"
	"bytes"
    "html/template"
    "sync"
    "time"

    "github.com/gofiber/fiber/v2"
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

    // Web server routes
    app.Get("/", homeHandler)
    app.Post("/start", startTaskHandler)

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

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        return c.Status(500).SendString("Error loading template: " + err.Error())
    }

    var renderedHTML bytes.Buffer
    if err := tmpl.Execute(&renderedHTML, struct{ Tasks []Task }{Tasks: tasks}); err != nil {
        return c.Status(500).SendString("Error rendering template: " + err.Error())
    }

    return c.Type("html").SendString(renderedHTML.String())
}


func startTaskHandler(c *fiber.Ctx) error {
    taskName := c.FormValue("task")

    mu.Lock()
    start := time.Now()
    mu.Unlock()

    mu.Lock()
    duration := time.Since(start)
    tasks = append(tasks, Task{Name: taskName, Duration: duration})
    mu.Unlock()

    return c.Redirect("/")
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
