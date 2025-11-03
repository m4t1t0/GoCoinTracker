# GoCoinTracker Architecture

GoCoinTracker is a Go application designed to track cryptocurrency assets. It follows best practices from the Go ecosystem, emphasizing simplicity, performance, and maintainability. The application leverages a PostgreSQL database for persistent storage of asset data.

## Core Principles

*   **Clean Architecture:** The application is structured to separate concerns, making it modular and testable. Business logic is independent of frameworks and databases.
*   **Modularity:** Components are designed to be small, focused, and reusable.
*   **Error Handling:** Go's idiomatic error handling (returning errors as the last return value) is consistently applied.
*   **Concurrency:** Goroutines and channels are utilized where appropriate for efficient handling of concurrent operations, especially for API requests and background tasks.
*   **Dependency Management:** Go Modules are used for managing project dependencies.

## Components

### 1. API Layer (`cmd/api`)

This layer handles incoming HTTP requests, routes them to appropriate handlers, and returns responses.

*   **`main.go`**: The entry point of the application, responsible for setting up the server and bootstrapping dependencies.
*   **`bootstrap/bootstrap.go`**: Manages the initialization of various application components, such as database connections and dependency injection.

### 2. Internal Logic (`internal`)

This package contains the core business logic and domain models, isolated from external concerns.

*   **`platform/server/server.go`**: Defines the HTTP server configuration and middleware.
*   **`platform/server/handler`**: Contains the HTTP handlers for different API endpoints.
    *   **`createAsset/createAsset.go`**: Handles the creation of new cryptocurrency assets.
    *   **`home/home.go`**: Provides a basic home endpoint, possibly for health checks or general information.

### 3. Database

PostgreSQL is used as the primary data store.

*   **Data Storage**: Stores information about cryptocurrency assets, user portfolios, and historical data.
*   **Schema Management**: Database migrations are managed to ensure schema evolution is smooth and controlled.

## Data Flow

1.  An HTTP request comes into the API layer (`cmd/api`).
2.  The request is routed to the appropriate handler within `internal/platform/server/handler`.
3.  The handler processes the request, potentially interacting with other internal services (e.g., a service layer for business logic, not explicitly shown but implied by clean architecture).
4.  Data persistence operations (e.g., saving a new asset) are performed against the PostgreSQL database.
5.  A response is constructed and sent back to the client.

## Development Environment

*   **Docker**: `Dockerfile` and `docker-compose.yml` are used to containerize the application and its dependencies (like PostgreSQL), ensuring a consistent development and deployment environment.
*   **Air**: `.air.toml` is configured for live-reloading during development, improving developer productivity.
*   **Makefile**: Provides common development tasks, such as running tests, building the application, and managing Docker containers.

This architecture aims to provide a robust, scalable, and maintainable foundation for GoCoinTracker.

## Development Guidelines

*   **Code Style**: The application follows the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) and [Effective Go](https://golang.org/doc/effective_go.html) style guides.
*   **Testing**: The application is tested using unit and integration tests.
*   **Documentation**: The application is documented using [GoDoc](https://godoc.org/github.com/gocointracker/gocointracker).
*   **Versioning**: The application uses [Semantic Versioning](https://semver.org/) for version control.
*   **Continuous Integration**: The application is continuously integrated using [GitHub Actions](https://github.com/gocointracker/gocointracker/actions).
*   **Continuous Delivery**: The application is continuously delivered using [GitHub Actions](https://github.com/gocointracker/gocointracker/actions).
*   **Simplicity First (KISS)**:
    - Always seek the simplest solution that meets the requirements.
    - When in doubt between two solutions, choose the simpler one.
    - Do not create abstractions "just in case" they are needed in the future.
*   **Pragmatic DRY**: Avoid significant duplication, but prioritize clarity over premature abstraction.
*   **WET Principle**: Write Everything Twice; do not introduce abstractions until you encounter two occurrences of the same code.
*   **Efficiency Without Premature Optimization**: Focus on clean design first, optimize only based on measurements.
*   **Minimize Dependencies**: Prefer standard language features over heavy frameworks.
*   **CUPID Properties**:
    - **C**ohesion: Modules have clear, focused responsibilities.
    - **U**nity: Components are independent and reusable.
    - **P**redictability: Consistent and easy-to-understand behavior.
    - **I**mmutability: Avoid mutable state whenever possible.
    - **D**ocumentation: Code is self-explanatory and well documented.

## For AIs and Automated Tools
This documentation is designed to be processed by AIs that assist in development:

* Follow Go recommendations and best practices
* Always respect the specified naming patterns
* Keep the principle of simplicity in mind when in doubt
* Avoid unnecessary abstractions until they are truly needed
