package storage

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID        int64
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
}

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
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
		return 0, err
	}
	return result.LastInsertId()
}

func GetActiveTask(db *sql.DB) (*Task, error) {
	var task Task
	err := db.QueryRow("SELECT id, name, start_time FROM tasks WHERE end_time IS NULL ORDER BY start_time DESC LIMIT 1").Scan(&task.ID, &task.Name, &task.StartTime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func StopTask(db *sql.DB, id int64) error {
	endTime := time.Now()
	_, err := db.Exec("UPDATE tasks SET end_time = ?, duration = ? WHERE id = ?", endTime, int(endTime.Sub(time.Now()).Seconds()), id)
	return err
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
		err := rows.Scan(&task.ID, &task.Name, &task.StartTime, &task.EndTime, &task.Duration)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
