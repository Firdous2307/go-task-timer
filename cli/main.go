/*
package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Firdous2307/go-task-timer/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Task struct {
	ID        int64
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

type CompletedTask struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Duration  float64   `json:"duration"`
}

var (
	mu sync.Mutex
	db *sql.DB
)

func main() {
	var err error
	db, err = storage.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		Views: nil,
	})

	app.Use(logger.New())
	app.Static("/", "../task-tracker-frontend/build")

	app.Post("/api/start", startTaskHandler)
	app.Post("/api/stop", stopTaskHandler)
	app.Get("/api/completed-tasks", completedTasksHandler)

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("../task-tracker-frontend/build/index.html")
	})

	// Start the Fiber app in a separate goroutine
	go func() {
		log.Fatal(app.Listen(":8080"))
	}()

	cliLoop() // Start the CLI loop
}

func cliLoop() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Start a task")
		fmt.Println("2. Stop a task")
		fmt.Println("3. View completed tasks")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input")
			continue
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			startTask()
		case "2":
			stopTask()
		case "3":
			viewCompletedTasks()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func startTask() {
	fmt.Print("Enter task name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		taskName := strings.TrimSpace(scanner.Text())
		if taskName == "" {
			fmt.Println("Task name cannot be empty.")
			return
		}

		id, err := storage.CreateTask(db, taskName)
		if err != nil {
			fmt.Println("Error starting task:", err)
			return
		}

		fmt.Printf("Task started: %s (ID: %d)\n", taskName, id)
		log.Printf("Task started via CLI: ID=%d, Name=%s", id, taskName)
	}
}

func stopTask() {
	fmt.Print("Enter task ID to stop: ")
	var taskID int64
	fmt.Scanln(&taskID)

	err := storage.StopTask(db, taskID)
	if err != nil {
		fmt.Println("Error stopping task:", err)
		return
	}

	// Retrieve the specific task that was stopped
	task, err := storage.GetTask(db, taskID)
	if err != nil {
		fmt.Println("Error retrieving task details:", err)
		return
	}

	duration := task.EndTime.Sub(task.StartTime)
	fmt.Printf("Task with ID %d stopped. Duration: %v\n", taskID, duration)
	log.Printf("Task stopped via CLI: ID=%d, Duration=%v", taskID, duration)
}

func viewCompletedTasks() {
	tasks, err := storage.GetCompletedTasks(db)
	if err != nil {
		fmt.Println("Error retrieving tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No completed tasks.")
		return
	}

	fmt.Println("Completed tasks:")
	for _, task := range tasks {
		fmt.Printf("- ID: %d, Name: %s, Duration: %v\n", task.ID, task.Name, task.Duration)
	}
}

func startTaskHandler(c *fiber.Ctx) error {
	taskName := c.FormValue("task")

	if taskName == "" {
		return c.Status(400).SendString("Task name is required.")
	}

	id, err := storage.CreateTask(db, taskName)
	if err != nil {
		return c.Status(500).SendString("Error starting task: " + err.Error())
	}

	return c.JSON(fiber.Map{"id": id, "message": "Task started"})
}

func stopTaskHandler(c *fiber.Ctx) error {
	taskID := c.FormValue("id")

	if taskID == "" {
		return c.Status(400).SendString("Task ID is required.")
	}

	id, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		return c.Status(400).SendString("Invalid task ID.")
	}

	err = storage.StopTask(db, id)
	if err != nil {
		return c.Status(500).SendString("Error stopping task: " + err.Error())
	}

	return c.SendString("Task stopped")
}

func completedTasksHandler(c *fiber.Ctx) error {
	tasks, err := storage.GetCompletedTasks(db)
	if err != nil {
		return c.Status(500).SendString("Error retrieving completed tasks: " + err.Error())
	}

	completedTasks := make([]CompletedTask, len(tasks))
	for i, task := range tasks {
		completedTasks[i] = CompletedTask{
			ID:        task.ID,
			Name:      task.Name,
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
			Duration:  task.Duration.Seconds(),
		}
	}

	return c.JSON(completedTasks)
}
*/
