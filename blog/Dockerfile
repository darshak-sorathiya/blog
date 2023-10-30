# Use the official golang 1.20 image as the base image
FROM golang:1.20 AS build

# Set the working directory in the container
WORKDIR /app

# Copy your Go source code to the container
COPY . .

# Build your Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o blog

# Use a lightweight base image with MySQL client
FROM alpine:3.14

# Install the MySQL client
RUN apk --no-cache add mysql-client

# Set the working directory in the container
WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=build /app/blog .

# Expose the port that your application will run on (if needed)
EXPOSE 8080

# Command to run your Go application
CMD ["./blog"]