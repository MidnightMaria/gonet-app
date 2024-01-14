# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

# Set the working directory
WORKDIR /app

# Copy the source code
COPY main.go .

# Run the application
CMD ["go", "run", "main.go"]
