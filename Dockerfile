# First stage: build the Go binary
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o wms-server cmd/server/main.go

RUN ls -la .

# Second stage: create a lightweight image to run the binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/wms-server .

# Expose port 8080 for WMS server
EXPOSE 8080

# Command to run the executable
CMD ["./wms-server"]