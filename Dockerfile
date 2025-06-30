FROM golang:1.24 AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
# Copy the entire application code
COPY . .
# Build the Go application
RUN go build -o myapp .
# Expose the port that your application will run on
EXPOSE 3000
# Command to run the executable
CMD ["./myapp"]