# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY *.go ./
COPY *.html ./
COPY *.css ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o terminal-beautified

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/terminal-beautified .
COPY --from=builder /app/*.html .
COPY --from=builder /app/*.css .

EXPOSE 8080

CMD ["./terminal-beautified"]
