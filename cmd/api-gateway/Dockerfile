FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем go mod файлы
COPY cmd/api-gateway/go.mod ./cmd/api-gateway/
COPY cmd/api-gateway/ ./cmd/api-gateway/

# Скачиваем зависимости
RUN cd cmd/api-gateway && go mod download

# Собираем приложение
RUN cd cmd/api-gateway && CGO_ENABLED=0 GOOS=linux go build -o /app/api-gateway .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]

