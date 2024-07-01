FROM golang:1.21-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o server ./cmd/server/main.go

FROM alpine:latest  
COPY --from=builder /app/server /app/server

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/app/server"]
