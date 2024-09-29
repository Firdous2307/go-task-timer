# Go-Task-Timer

This is a simple web-based task timer application written in Go. It allows users to track the duration of tasks, offering an intuitive web interface to start a task, stop it, and view the task's duration.

## Features

- **Start a Task**: Input a task name to track its duration.
- **View Task Durations**: Display completed tasks and their durations.
- **Lightweight Web Server**: Powered by Go's `net/http` package for quick hosting.
- **Responsive Design**: User-friendly interface for task management.
- **Command-Line Interface**: Allows users to start and manage tasks directly from the terminal.

## Web Interface:

Start Task: Allows users to start a new task and track its duration.

View Completed Tasks: Displays completed tasks along with the time spent on each.

## Command-Line Interface

- **CLI Functionality**: Users can interact with the application using a command-line interface to start, view, and exit tasks.

## File Structure

- `main.go`: Server logic and task tracking functionality.
- `templates/index.html`: Basic HTML template for the task interface.
- `.github/workflows/go-ci.yml`: GitHub Actions workflow for testing and CI.
