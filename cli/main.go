package main

import (
    "log"
    "fmt"
    "sync"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/template/html/v2"

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
    engine := html.New("./templates", ".html")

    engine.Reload(true) // Reload templates on each render
    engine.Debug(true)  // Print parsed templates for debugging

    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Use(logger.New())
    app.Static("/static", "./static")

    app.Get("/", homeHandler)
    app.Post("/start", startTaskHandler)
    app.Post("/stop", stopTaskHandler)

    // Start the Fiber app in a separate goroutine
    go func() {
        log.Fatal(app.Listen(":8080"))
    }()

    cliLoop() // Start the CLI loop
}

// CLI Loop for user interaction
func cliLoop() {
    cliTasks := make(map[string]time.Duration)

    for {
        // Display options to the user
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
            startTask(cliTasks)
        case 2:
            viewTasks(cliTasks)
        case 3:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func startTask(cliTasks map[string]time.Duration) {
    var taskName string
    fmt.Print("Enter task name: ")
    fmt.Scan(&taskName)

    // Example of starting a task
    cliTasks[taskName] = 0 // Initialize task duration
    fmt.Println("Task started:", taskName)
}

func viewTasks(cliTasks map[string]time.Duration) {
    fmt.Println("Current tasks:")
    for name, duration := range cliTasks {
        fmt.Printf("Task: %s, Duration: %v\n", name, duration)
    }
}

func homeHandler(c *fiber.Ctx) error {
    mu.Lock()
    defer mu.Unlock()

    return c.Render("index", fiber.Map{"Tasks": tasks})
}

func startTaskHandler(c *fiber.Ctx) error {
    taskName := c.FormValue("task")
    if taskName == "" {
        return c.Status(400).SendString("Task name is required.")
    }

    mu.Lock()
    tasks = append(tasks, Task{Name: taskName, Duration: 0})
    mu.Unlock()

    return c.SendString("Task started")
}

func stopTaskHandler(c *fiber.Ctx) error {
    taskName := c.FormValue("task")
    durationStr := c.FormValue("duration")
    if taskName == "" || durationStr == "" {
        return c.Status(400).SendString("Task name and duration are required.")
    }

    duration, err := time.ParseDuration(durationStr + "s")
    if err != nil {
        return c.Status(400).SendString("Invalid duration.")
    }

    mu.Lock()
    for i, task := range tasks {
        if task.Name == taskName && task.Duration == 0 {
            tasks[i].Duration = duration
            break
        }
    }
    mu.Unlock()

    return c.SendString("Task stopped")
}
