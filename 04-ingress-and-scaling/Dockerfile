# Use docker multistage builds, name build container; builder
FROM golang:alpine as builder

# Golang build flags
ENV CGO_ENABLED=0
ENV GOOS=linux

# Create directory for app
RUN mkdir /app

# Set working directory to /app
WORKDIR /app

# Copy the main.go file to our working directory /app
COPY main.go /app

# Build our golang application
RUN go build -o server /app/main.go

# "Production" image, use scratch base image
FROM scratch

# Copy the binary from the build container
COPY --from=builder app/server ./server

# Run the built and copied binary
CMD ["./server"]