# Containerization with Docker

## Why Containerization?

1. **Consistency**: Ensures application runs identically across all environments
2. **Isolation**: Packages all dependencies within the container
3. **Scalability**: Easy horizontal scaling and orchestration
4. **DevOps Integration**: Seamless CI/CD pipeline integration
5. **Resource Efficiency**: Lightweight compared to traditional VMs

## Getting Started

1. Build the image:
   ```bash
   docker build -t go-task-timer .
   ```

2. Run the container:
   ```bash
   docker-compose up
   ```