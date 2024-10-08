# Use the official Golang image as a base
FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Command to run the executable
CMD ["make", "run"]
