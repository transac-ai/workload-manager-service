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

# Build the Go binary with static linking to avoid dependencies
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o wms-server ./core/server.go


RUN ls -la .

# Second stage: create a lightweight image to run the binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/wms-server .

# Set GO_ENV to production
ENV GO_ENV=production

# Expose port 8080 for WMS server
EXPOSE 8080

# Command to run the executable
CMD ["./wms-server"]