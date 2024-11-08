# Use Go 1.21
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (if they exist)
COPY go.mod go.sum* ./

# Download and install dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Expose the port
EXPOSE 8000

# Run the executable
CMD ["./main"]
