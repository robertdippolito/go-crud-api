# Use the official Go image as the base
FROM golang:1.20
# Set the working directory
WORKDIR /usr/src/app
# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download
# Copy the entire application code
COPY . .
# Build the Go application
RUN go build -o myapp main.go
# Expose the port that your application will run on
EXPOSE 3000
# Command to run the executable
CMD ["./myapp"]