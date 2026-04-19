# Stage 1: Build
FROM golang:1.25-alpine AS build

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o ticket .

# Stage 2: Run
FROM alpine:latest

WORKDIR /tixi

COPY --from=build /app/ticket .

EXPOSE 8081

CMD ["./ticket"]