# Use a Go base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /go/src/github.com/Nachofra/final-esp-backend-3

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all the source code from the relative location "cmd/api/" to the current working directory
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/
COPY ./docs/ ./docs/

# Run go mod tidy
RUN go mod tidy

# Compile the application
RUN go build -o main ./cmd/api/

# Expose the port on which the application will listen
EXPOSE 8080

# Command to run the application when the container starts
CMD ["./main"]
