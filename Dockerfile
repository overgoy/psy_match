# Stage 1: сборка приложения
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o psymatch ./cmd/main.go

# Stage 2: создание минимального образа
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/psymatch .
COPY ./config /app/config
COPY ./migrations /app/migrations
EXPOSE 8080
CMD ["./psymatch"]
