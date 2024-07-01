FROM golang:1.21-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o client ./cmd/client/main.go

FROM alpine:latest  
COPY --from=builder /app/client /app/client

# Command to run the executable
CMD ["/app/client"]
