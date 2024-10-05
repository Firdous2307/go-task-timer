package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

type CompletedTask struct {
	Name     string  `json:"Name"`
	Duration float64 `json:"Duration"`
}

var (
	tasks          []Task
	mu             sync.Mutex
	completedTasks []CompletedTask
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
	app.Get("/completed-tasks", completedTasksHandler)

	// Start the Fiber app in a separate goroutine
	go func() {
		log.Fatal(app.Listen(":8080"))
	}()

	cliLoop() // Start the CLI loop
}

// CLI Loop for user interaction
func cliLoop() {
	cliTasks := make(map[string]time.Time)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Start a task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input")
			continue
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			startTask(cliTasks)
		case "2":
			viewTasks(cliTasks)
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func startTask(cliTasks map[string]time.Time) {
	var taskName string
	var duration int

	fmt.Print("Enter task name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		taskName = strings.TrimSpace(scanner.Text())
	}
	if taskName == "" {
		fmt.Println("Task name cannot be empty.")
		return
	}

	fmt.Print("Enter task duration in seconds (0 for no limit): ")
	_, err := fmt.Scanf("%d\n", &duration)
	if err != nil {
		fmt.Println("Invalid duration. Please enter a number.")
		return
	}

	if _, exists := cliTasks[taskName]; exists {
		fmt.Println("Task already exists and is running.")
		return
	}

	cliTasks[taskName] = time.Now()
	fmt.Println("Task started:", taskName)

	if duration > 0 {
		go func() {
			time.Sleep(time.Duration(duration) * time.Second)
			if _, exists := cliTasks[taskName]; exists {
				delete(cliTasks, taskName)
				fmt.Printf("Task '%s' automatically stopped after %d seconds.\n", taskName, duration)
			}
		}()
	}
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

	// Add the completed task to the slice
	completedTasks = append(completedTasks, CompletedTask{
		Name:     taskName,
		Duration: duration.Seconds(),
	})
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
	durationStr := c.FormValue("duration")

	if taskName == "" {
		return c.Status(400).SendString("Task name is required.")
	}

	duration, err := time.ParseDuration(durationStr + "s")
	if err != nil {
		return c.Status(400).SendString("Invalid duration format.")
	}

	mu.Lock()
	tasks = append(tasks, Task{Name: taskName, Duration: duration})
	mu.Unlock()

	// Start a goroutine to handle task completion
	go handleTaskCompletion(taskName, duration)
	return c.SendString("Task started")
}

func handleTaskCompletion(taskName string, duration time.Duration) {
	time.Sleep(duration)

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.Name == taskName {
			// Remove the task from the slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			// Add the task to completedTasks
			completedTasks = append(completedTasks, CompletedTask{
				Name:     taskName,
				Duration: duration.Seconds(),
			})
			// Print the completed task
			fmt.Printf("Task completed: %s (Duration: %v)\n", taskName, duration)
			break
		}
	}
}

func completedTasksHandler(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	return c.JSON(completedTasks)
}

func stopTaskHandler(c *fiber.Ctx) error {
	taskName := c.FormValue("task")
	durationStr := c.FormValue("duration")

	if taskName == "" {
		return c.Status(400).SendString("Task name is required.")
	}

	duration, err := time.ParseDuration(durationStr + "s")
	if err != nil {
		return c.Status(400).SendString("Invalid duration format.")
	}

	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.Name == taskName {
			tasks[i].Duration = duration
			completedTasks = append(completedTasks, CompletedTask{Name: taskName, Duration: duration.Seconds()})
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.SendString("Done")
		}
	}

	return c.SendString("Done")
}
