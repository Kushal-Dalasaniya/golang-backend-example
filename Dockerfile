# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git (required for go mod in some cases)
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
