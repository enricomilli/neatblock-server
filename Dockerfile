
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache make git

# Copy only the files needed for dependency installation
COPY go.mod go.sum ./
COPY Makefile ./

# Install dependencies
RUN make install

# Copy the rest of the application code
COPY . .

# Build the application
RUN make build

# Start a new stage for a smaller final image
FROM alpine:latest

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy the binary and necessary files from the builder stage
COPY --from=builder /app/bin/server .

# Create a non-root user
RUN adduser -D appuser
USER appuser

EXPOSE 8080

CMD ["./server"]
