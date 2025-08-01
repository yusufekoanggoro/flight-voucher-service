FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Copy go.mod dan go.sum lalu download dependencies
COPY go.mod go.sum ./

RUN go mod tidy && go mod download

COPY . .

RUN mkdir -p /app/data

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]