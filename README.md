# Go-Task-Timer

Go-Task-Timer is a versatile task timing application written in Go. It provides both a web interface and a command-line interface for users to track and manage their tasks efficiently.

## Features

- **Dual Interface**: 
  - Web-based interface for easy access through a browser
  - Command-line interface for quick task management from the terminal
- **Task Management**:
  - Start and stop tasks
  - View completed tasks with their durations
- **Concurrent Operation**: Run both web and CLI interfaces simultaneously

## Web Interface

The web interface is built using the Fiber web framework, offering a responsive and user-friendly experience:

- Start a new task
- Stop the current task
- View a list of completed tasks with their durations

## Command-Line Interface (CLI)

The CLI provides a quick way to manage tasks directly from the terminal:

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

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Execute `go run cli/main.go` to start the application
4. Access the web interface at `http://localhost:8080`
5. Use the CLI by following the on-screen prompts

## Testing

Run the tests using the command: `go test ./tests`

## Continuous Integration

The project uses GitHub Actions for continuous integration, ensuring code quality and test coverage with each push to the repository.
