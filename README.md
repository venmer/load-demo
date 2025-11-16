# Load Demo - Демо проект для воркшопа по нагрузочному тестированию

Этот проект демонстрирует простую микросервисную архитектуру на Go

## Архитектура

Проект состоит из следующих компонентов:

1. **User Service** (`cmd/user-service`) - микросервис для управления пользователями (CRUD операции)
2. **API Gateway** (`cmd/api-gateway`) - шлюз для маршрутизации запросов к микросервисам
3. **PostgreSQL** - база данных для хранения данных пользователей

## Требования

- Docker и Docker Compose
- Go 1.24+ (для локальной разработки)

## Быстрый старт

1. **Клонируйте репозиторий и перейдите в директорию проекта:**

```bash
cd load-demo
```

2. **Создайте файл `.env` со следующими переменными:**

```bash
POSTGRES_USER=
DB_PASSWORD=
DB_NAME=

GF_SECURITY_ADMIN_USER=
GF_SECURITY_ADMIN_PASSWORD=
```
заполните переменные значениями

3. **Запустите все сервисы через Docker Compose:**

```bash
make app-up
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

## Остановка сервисов

```bash
docker-compose down
```

Для полной очистки (включая volumes):

```bash
make app-stop
```

## Воркшоп

### Метрики нагрузочного тестирования
Эта часть включает в себя следующее:
- Разбор конфигурации демо сервиса
- Разбор конфигурации мониторинга
- Запуск сервисов мониторинга
- Импорт Grafana дашборда
- Настройка Grafana dashboard
- Рассмотрение типов метрик
- Взаимодействие с сервисом и отслеживание метрик

Заметки:

запуск/проверка доступности сервисов мониторинга:
```bash
make mon-up
```

Перейти в Grafana
```
http://localhost:3000
```
- добавить DS к viktoria-metrics
```bash
http://victoria-metrics:8428
```


### Нагрузочное тестирование

- Выбор профиля нагрузки
- Выбор инструмента (k6, locust)
- Подготовка дашборда для мониторинга
- Запуск первого теста и отладка
- Составление профиля нагрузки
- Запуска теста и анализ результата

Заметки:

К профилю нагрузки нужно заранее определить требования 