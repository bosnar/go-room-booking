# Go Room Booking API

This project is a backend service for managing meeting room bookings, built using the Go programming language. The service follows Clean Architecture principles to ensure modularity and testability. The Gin framework is used as the HTTP server for handling API requests.

## Architecture

The project structure is organized as follows:

- **cmd/**: Contains the main application entry point.
- **config/**: Stores configuration files and logic for the application.
- **controller/**: Handles incoming HTTP requests and maps them to use cases.
- **domain/**: Defines the core business logic and models for the application.
- **repository/**: Manages data persistence and retrieval from the data store.
- **router/**: Configures routing for all API endpoints using the Gin framework.
- **test/**: Contains unit and integration tests.
- **usecase/**: Encapsulates the business logic and orchestrates interaction between repositories and controllers.

## Features

- RESTful API using Gin
- Clean Architecture for scalable and maintainable code
- Test coverage for critical components

## Requirements

- Go 1.20+
