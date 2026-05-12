# vps-dash-server

A lightweight Go HTTP API that exposes system metrics (CPU, memory, processes, and historical snapshots) for VPS dashboard frontends.

## Features

- **CPU & Memory usage** — real-time or cached endpoints
- **Process list** — list all running processes or query by PID
- **Rolling snapshots** — configurable historical metrics buffer
- **Token-based auth** — simple Bearer token authentication
- **Scheduled caching** — background updates to reduce API latency

## Requirements

- Go 1.22.5+
- Linux/macOS/Windows (uses [`gopsutil`](https://github.com/shirou/gopsutil))

## Quick start

```bash
# Clone the repository
git clone https://github.com/JaroxCraft/vps-dash-server.git
cd vps-dash-server

# Copy the example environment file and edit it
cp example.env .env

# Run the server
go run .
```

## Configuration

All configuration is done via environment variables (or a `.env` file thanks to [`godotenv`](https://github.com/joho/godotenv)).

| Variable            | Default      | Description                              |
|---------------------|--------------|------------------------------------------|
| `AUTH_TOKEN`        | *(required)* | Bearer token for API authentication      |
| `API_PORT`          | `8080`       | HTTP server port                         |
| `SNAPSHOT_INTERVAL` | `15s`        | How often to record a metrics snapshot   |
| `CACHE_INTERVAL`    | `5s`         | How often to refresh the usage cache     |

> ⚠️ **Security note:** `AUTH_TOKEN` must be set. The server will refuse to start without it.

## API

All endpoints require an `Authorization: Bearer <AUTH_TOKEN>` header.

### Stats

| Method | Endpoint             | Query params | Description                              |
|--------|----------------------|--------------|------------------------------------------|
| `GET`  | `/stats/cpu/`        | `?cache=true\|false` | CPU usage percentage (0–100)    |
| `GET`  | `/stats/memory/`     | `?cache=true\|false` | Memory usage percentage (0–100) |

### Processes

| Method | Endpoint             | Description                              |
|--------|----------------------|------------------------------------------|
| `GET`  | `/processes/`        | List all running processes               |
| `GET`  | `/processes/{pid}/`  | Get a single process by PID              |

### Snapshots

| Method | Endpoint             | Description                              |
|--------|----------------------|------------------------------------------|
| `GET`  | `/snapshots/`        | Rolling buffer of the last 8 snapshots   |

## Build

```bash
go build -o vps-dash-server .
```

## Test

```bash
go test ./...
```

## License

MIT
