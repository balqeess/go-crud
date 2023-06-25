# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download 

RUN go get github.com/githubnemo/CompileDaemon

# Copy the rest of the project files
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the desired port (e.g., 8080)
EXPOSE 3000

# Set the command to run the executable
CMD ["./main"]

