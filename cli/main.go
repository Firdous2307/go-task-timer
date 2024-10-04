package main

import (
	"fmt"
	"log"
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
	// Updated paths in main.go
	engine := html.New("../web/templates", ".html") // Template path relative to cli foldr
	engine.Reload(true)                             // Reload templates on each render
	engine.Debug(true)                              // Print parsed templates for debugging

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())
	app.Static("/static", "../web/static") // Static path relative to cli folder

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
	cliTasks := make(map[string]time.Time)

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

func startTask(cliTasks map[string]time.Time) {
	var taskName string
	fmt.Print("Enter task name: ")
	fmt.Scan(&taskName)
	if taskName == "" {
		fmt.Println("Task name cannot be empty.")
		return
	}

	if _, exists := cliTasks[taskName]; exists {
		fmt.Println("Task already exists and is running.")
		return
	}

	cliTasks[taskName] = time.Now()
	fmt.Println("Task started:", taskName)
}

func stopTask(cliTasks map[string]time.Time) {
	var taskName string
	fmt.Print("Enter task name to stop: ")
	fmt.Scanln(&taskName)

	startTime, exists := cliTasks[taskName]
	if !exists {
		fmt.Println("Task not found or already stopped.")
		return
	}

	duration := time.Since(startTime)
	delete(cliTasks, taskName)
	fmt.Printf("Task '%s' stopped. Duration: %v\n", taskName, duration.Round(time.Second))
}

func viewTasks(cliTasks map[string]time.Time) {
	if len(cliTasks) == 0 {
		fmt.Println("No active tasks.")
		return
	}

	fmt.Println("Active tasks:")
	for name, startTime := range cliTasks {
		duration := time.Since(startTime)
		fmt.Printf("- %s (running for %v)\n", name, duration.Round(time.Second))
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
	if taskName == "" {
		return c.Status(400).SendString("Task name is required.")
	}

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.Name == taskName && task.Duration == 0 {
			tasks[i].Duration = time.Since(time.Now()) // Calculate duration
			return c.SendString("Task stopped")
		}
	}

	return c.Status(400).SendString("Task not found or already stopped.")
}
