package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Firdous2307/go-task-timer/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
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
		Views: html.New("./web/templates", ".html"),
	})

	app.Use(logger.New())
	app.Static("/static", "./web/static")

	app.Get("/", indexHandler)
	app.Post("/start", startTaskHandler)
	app.Post("/stop", stopTaskHandler)

	log.Fatal(app.Listen(":8080"))
}

func indexHandler(c *fiber.Ctx) error {
	tasks, err := storage.GetCompletedTasks(db)
	if err != nil {
		return c.Status(500).SendString("Error retrieving completed tasks")
	}

	return c.Render("index", fiber.Map{
		"Tasks": tasks,
	})
}

func startTaskHandler(c *fiber.Ctx) error {
	taskName := c.FormValue("task")

	if taskName == "" {
		return c.Redirect("/")
	}

	_, err := storage.CreateTask(db, taskName)
	if err != nil {
		return c.Status(500).SendString("Error starting task: " + err.Error())
	}

	return c.Redirect("/")
}

func stopTaskHandler(c *fiber.Ctx) error {
	tasks, err := storage.GetActiveTasks(db)
	if err != nil {
		return c.Status(500).SendString("Error retrieving active tasks")
	}

	if len(tasks) > 0 {
		err = storage.StopTask(db, tasks[0].ID)
		if err != nil {
			return c.Status(500).SendString("Error stopping task: " + err.Error())
		}
	}

	return c.Redirect("/")
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
