# Continuous Integration and Deployment

Go-Task-Timer uses GitHub Actions for continuous integration and deployment. This ensures code quality, test coverage, and security with each push to the repository.

The CI/CD pipeline includes:
- Running all unit tests
- Building the application
- Performing a security scan
- Deploying to production (only for pushes to the main branch)

When working in a branch or on a pull request, the pipeline includes:
- Running all unit tests
- Building the application
- Performing a security scan

[Learn how to contribute](contributing.md)