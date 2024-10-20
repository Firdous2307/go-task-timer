package storage

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite" // Updated to use modernc.org/sqlite driver
)

type Task struct {
	ID        int64
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./tasks.db") // Change to "sqlite"
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		start_time DATETIME,
		end_time DATETIME,
		duration INTEGER
	)`)

	return db, err
}

func CreateTask(db *sql.DB, name string) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (name, start_time) VALUES (?, ?)", name, time.Now())
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	log.Printf("Task created: ID=%d, Name=%s", id, name)
	return id, nil
}

func StopTask(db *sql.DB, id int64) error {
	endTime := time.Now()
	var startTime time.Time
	err := db.QueryRow("SELECT start_time FROM tasks WHERE id = ?", id).Scan(&startTime)
	if err != nil {
		log.Printf("Error retrieving start time: %v", err)
		return err
	}
	duration := endTime.Sub(startTime)
	_, err = db.Exec("UPDATE tasks SET end_time = ?, duration = ? WHERE id = ?",
		endTime,
		int(duration.Seconds()),
		id)
	if err != nil {
		log.Printf("Error stopping task: %v", err)
		return err
	}
	log.Printf("Task stopped: ID=%d, Duration=%v", id, duration)
	return nil
}

func GetTask(db *sql.DB, id int64) (*Task, error) {
	var task Task
	err := db.QueryRow("SELECT id, name, start_time, end_time, duration FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Name, &task.StartTime, &task.EndTime, &task.Duration)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetCompletedTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, name, start_time, end_time, duration FROM tasks WHERE end_time IS NOT NULL ORDER BY end_time DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		var durationSeconds int64
		err := rows.Scan(&task.ID, &task.Name, &task.StartTime, &task.EndTime, &durationSeconds)
		if err != nil {
			return nil, err
		}
		task.Duration = time.Duration(durationSeconds) * time.Second
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetActiveTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, name, start_time FROM tasks WHERE end_time IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name, &t.StartTime)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
