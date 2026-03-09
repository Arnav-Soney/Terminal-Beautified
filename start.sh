#!/bin/bash

# Terminal Beautified - Startup Script

echo "🚀 Starting Terminal Beautified Server..."

# Set default port if not specified
export PORT=${PORT:-8080}

# Run the Go server with all source files
go run *.go
