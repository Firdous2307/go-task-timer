---
layout: default
title: Go-Task-Timer Documentation
---

# Go-Task-Timer

Go-Task-Timer is a versatile task timing application written in Go. It provides both a web interface and a command-line interface for users to track and manage their tasks efficiently.

## Features

- **Dual Interface**: 
  - Web-based interface for easy access through a browser
  - Command-line interface for quick task management from the terminal
- **Task Management**:
  - Start and stop tasks
  - View completed tasks with their durations
  - Tasks are stored in memory during runtime
- **Concurrent Operation**: Run both web and CLI interfaces simultaneously

## Installation

1. Ensure you have Go installed on your system
2. Clone the repository:
   ```
   git clone https://github.com/yourusername/go-task-timer.git
   ```
3. Navigate to the project directory:
   ```
   cd go-task-timer
   ```
4. Install dependencies:
   ```
   go mod tidy
   ```

## Usage

### Starting the Application

Run the following command in the project root:
```go run main.go task.go```
This will start both the web interface and the CLI.


### Web Interface

Access the web interface by opening a browser and navigating to `http://localhost:8080`. Here you can:

- Start a new task
- Stop the current task

### Command-Line Interface (CLI)

The CLI provides a quick way to manage tasks directly from the terminal. Follow the on-screen prompts to:

1. Start a task
2. View completed tasks
3. Exit the application

## Project Structure

- `cli/main.go`: Main entry point, sets up both web and CLI interfaces
- `internal/task/tasks.go`: Core task management logic
- `templates/`: HTML templates for the web interface
- `static/`: Static assets for the web interface
- `tests/main_test.go`: Unit tests for the application

## Technologies Used

- Go programming language
- Fiber web framework
- HTML templates
- Goroutines for concurrent operation

## Testing

Run the tests using the command:

```go test ./tests```


## Continuous Integration

The project uses GitHub Actions for continuous integration, ensuring code quality and test coverage with each push to the repository.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

