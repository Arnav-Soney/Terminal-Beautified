# Backend-Driven Configuration

## Overview

The frontend now loads all constants from the Go backend via `/api/config` endpoint. This allows you to update configuration without redeploying the frontend.

## Architecture

- **Backend**: Go HTTP server serves the HTML/CSS/JS files and provides a REST API
- **Frontend**: Fetches configuration from `/api/config` on page load
- **Fallback**: If backend is unavailable, frontend uses hardcoded fallback values

## Running the Server

```bash
# Development
go run main.go

# Build
go build -o terminal-beautified main.go

# Run built binary
./terminal-beautified
```

Server runs on `http://localhost:8080` by default.

## Updating Configuration

To change the constants shown in the frontend:

1. Edit the `getConfig()` function in [main.go](../main.go)
2. Restart the server
3. Refresh the browser - new values load automatically

No frontend redeployment needed! ✨

## API Endpoints

- `GET /` - Serves the HTML page
- `GET /style.css` - Serves the CSS file
- `GET /api/config` - Returns JSON configuration

## Environment Variables

- `PORT` - Server port (default: 8080)

## Deployment

For production deployment:

1. Build the binary: `go build -o app main.go`
2. Set the `PORT` environment variable if needed
3. Run the binary: `./app`

The server will serve the static files from the `Terminal-Beautified` directory and provide the configuration API.
