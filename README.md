# Load Demo - Демо проект для демонстрации нагрузочного тестирования

Этот проект демонстрирует простую микросервисную архитектуру на Go

## Архитектура

Проект состоит из следующих компонентов:

1. **User Service** (`cmd/user-service`) - микросервис для управления пользователями (CRUD операции)
2. **API Gateway** (`cmd/api-gateway`) - шлюз для маршрутизации запросов к микросервисам
3. **PostgreSQL** - база данных для хранения данных пользователей

## Структура проекта

```
load-demo/
├── cmd/
│   ├── user-service/       # Микросервис пользователей
│   │   ├── main.go
│   │   ├── internal/
│   │   │   ├── database/   # Подключение к БД
│   │   │   ├── handlers/   # HTTP handlers
│   │   │   └── models/     # Модели данных
│   │   ├── Dockerfile
│   │   └── go.mod
│   └── api-gateway/        # API Gateway
│       ├── main.go
│       ├── Dockerfile
│       └── go.mod
├── migrations/              # SQL миграции
├── docker-compose.yml       # Оркестрация сервисов
└── README.md
```

## Требования

- Docker и Docker Compose
- Go 1.21+ (для локальной разработки)

## Быстрый старт

1. **Клонируйте репозиторий и перейдите в директорию проекта:**

```bash
cd load-demo
```

2. **Создайте файл `.env` на основе `.env.example`:**

```bash
cp .env.example .env
```

Отредактируйте `.env` при необходимости.

3. **Запустите все сервисы через Docker Compose:**

```bash
docker-compose up --build
```

4. **Проверьте работу сервисов:**

- API Gateway: http://localhost:8080
- User Service: http://localhost:8081
- PostgreSQL: localhost:5432

## API Endpoints

### Через API Gateway (порт 8080)

- `GET /health` - проверка здоровья сервиса
- `GET /api/users` - получить список всех пользователей
- `GET /api/users/{id}` - получить пользователя по ID
- `POST /api/users` - создать нового пользователя
- `PUT /api/users/{id}` - обновить пользователя
- `DELETE /api/users/{id}` - удалить пользователя

### Прямой доступ к User Service (порт 8081)

Те же endpoints доступны напрямую через порт 8081.

## Примеры использования

### Создать пользователя

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Иван Иванов", "email": "ivan@example.com"}'
```

### Получить всех пользователей

```bash
curl http://localhost:8080/api/users
```

### Получить пользователя по ID

```bash
curl http://localhost:8080/api/users/1
```

### Обновить пользователя

```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Иван Петров", "email": "ivan.petrov@example.com"}'
```

### Удалить пользователя

```bash
curl -X DELETE http://localhost:8080/api/users/1
```

## Локальная разработка

Для разработки без Docker:

1. **Запустите PostgreSQL:**

```bash
docker run -d \
  --name postgres-dev \
  -e POSTGRES_USER=demo_user \
  -e POSTGRES_PASSWORD=demo_password \
  -e POSTGRES_DB=demo_db \
  -p 5432:5432 \
  postgres:15-alpine
```

2. **Запустите User Service:**

```bash
cd cmd/user-service
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=demo_user
export DB_PASSWORD=demo_password
export DB_NAME=demo_db
export SERVER_PORT=8081
go run main.go
```

3. **Запустите API Gateway:**

```bash
cd cmd/api-gateway
export USER_SERVICE_URL=http://localhost:8081
export SERVER_PORT=8080
go run main.go
```

## Переменные окружения

### User Service

- `DB_HOST` - хост PostgreSQL (по умолчанию: localhost)
- `DB_PORT` - порт PostgreSQL (по умолчанию: 5432)
- `DB_USER` - пользователь БД (по умолчанию: demo_user)
- `DB_PASSWORD` - пароль БД (по умолчанию: demo_password)
- `DB_NAME` - имя БД (по умолчанию: demo_db)
- `SERVER_PORT` - порт сервиса (по умолчанию: 8081)

### API Gateway

- `USER_SERVICE_URL` - URL User Service (по умолчанию: http://localhost:8081)
- `SERVER_PORT` - порт шлюза (по умолчанию: 8080)

## Остановка сервисов

```bash
docker-compose down
```

Для полной очистки (включая volumes):

```bash
docker-compose down -v
```
