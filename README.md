# GoCoinTracker

A small, personal project written in Go to explore the Go ecosystem, best practices, and language idioms. This repository is intended for learning purposes only — it is not a production-ready application.

## Goals

- Practice clean, simple Go code and project structure
- Experiment with HTTP APIs using Fiber
- Learn validation, routing, and basic handler patterns
- Keep things minimal and easy to understand

## What it does (today)

- Starts an HTTP server (Fiber) on a configurable port
- Exposes a basic health/home endpoint
- Provides a sample endpoint to validate and echo asset creation payloads

## Status

Early stage and unstable. Breaking changes are expected. Features are incomplete and may change or be removed.

## Quick Start

### Option 1: Run with Docker (recommended)

Prerequisites:
- Docker and Docker Compose
- An external Docker network named `quadrant-network` (create it once if you don't already have it):

```bash
docker network create quadrant-network
```

Commands:

```bash
# Build the image
make build

# Start the service
timeout 1s bash -c ': >/dev/tcp/127.0.0.1/3000' 2>/dev/null || true
make start

# Stop the service
make down
```

This will start the API on http://localhost:3000 with `HTTP_PORT=3000`.

Useful helper:

```bash
# Open a shell in the running container
make sh
```

### Option 2: Run locally (without Docker)

Prerequisites:
- Go toolchain installed

Run:

```bash
export HTTP_PORT=3000
go run ./cmd/api
```

The server will listen on `:3000`.

## API Overview

Base URL: `http://localhost:3000`

- GET `/`
  - Health/home check
  - Response example:
    ```json
    {"status": "ok", "version": "0.1"}
    ```

- POST `/api/v1/assets`
  - Accepts JSON body, validates fields, and echoes them back on success
  - Request body:
    ```json
    {
      "asset": "BTC",
      "interval": 60
    }
    ```
    - `asset`: required, alphanumeric, 3–50 characters
    - `interval`: required, integer > 0
  - Success (200) response:
    ```json
    {
      "asset": "BTC",
      "interval": 60
    }
    ```
  - Validation errors return 400 with details

### cURL examples

```bash
# Health
curl -s http://localhost:3000/

# Create asset (valid)
curl -s -X POST http://localhost:3000/api/v1/assets \
  -H 'Content-Type: application/json' \
  -d '{"asset":"BTC","interval":60}'

# Create asset (invalid: too short)
curl -s -X POST http://localhost:3000/api/v1/assets \
  -H 'Content-Type: application/json' \
  -d '{"asset":"BT","interval":60}'
```

## Configuration

- `HTTP_PORT` — the TCP port for the HTTP server (e.g., `3000`)
- `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_HOST`, `POSTGRES_DB` — database connection settings used to build the DSN `postgres://user:pass@host/db?sslmode=disable`.
- `GORM_LOG_LEVEL` — optional; set to `info` locally for verbose SQL logs. Defaults to `warn` if unset.

Example (local):
```bash
export HTTP_PORT=3000
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export POSTGRES_HOST=localhost:5432
export POSTGRES_DB=gocointracker
# Optional, for verbose SQL output in dev
export GORM_LOG_LEVEL=info
```

## Database & Migrations

- Schema is managed exclusively via SQL migrations located in `migrations/`.
- There is no GORM `AutoMigrate` path enabled in any environment.
- Use the provided Makefile tasks to manage your database:
  - `make db-create` — create the database (if not exists)
  - `make db-migrate` — apply all up migrations
  - `make db-drop` — drop the database
  - `make db-fresh` — drop, create, and migrate

## GORM Logging

- Default GORM log level is `warn`.
- For local debugging, set `GORM_LOG_LEVEL=info` to see executed SQL statements and timings.

## Project Layout

- `cmd/api` — entrypoint and bootstrap
- `internal/platform/server` — Fiber server and route registration
  - `handler/home` — GET `/`
  - `handler/createAsset` — POST `/api/v1/assets`
  - `docs/architecture` — additional notes about the intended architecture; see [Architecture Guidelines](docs/architecture/index.md)
  - `Dockerfile`, `docker-compose.yml`, `Makefile` — development tooling

## Notes & Limitations

- Educational project: not audited, not optimized, not production-ready
- Interfaces, error handling, and structure are intentionally simple to favor clarity over features

## License

No license specified. If you plan to use or extend this code, please consult the repository owner.
