# Этап сборки
FROM golang:1.22-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd

# Финальный образ
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .              
COPY config config/    
COPY certs/ ./certs/ 

CMD ["./main"]
