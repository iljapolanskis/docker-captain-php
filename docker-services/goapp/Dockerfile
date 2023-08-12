# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files to the container
COPY ../../app/main.go ./
COPY ../../app/go.mod ./

# Download all dependencies
# RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Command to run the application
CMD ["./main"]
