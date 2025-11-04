FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем go mod файлы
COPY cmd/user-service/go.mod ./cmd/user-service/
COPY cmd/user-service/ ./cmd/user-service/

# Скачиваем зависимости
RUN cd cmd/user-service && go mod download

# Собираем приложение
RUN cd cmd/user-service && CGO_ENABLED=0 GOOS=linux go build -o /app/user-service .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/user-service .

EXPOSE 8081

CMD ["./user-service"]

