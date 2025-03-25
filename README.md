# User API Go

API для управления пользователями на Go

## Возможности

- RESTful API для управления пользователями
- Контейнеризация с помощью Docker
- Миграции базы данных
- Тесты

## Структура проекта

```
user-api-go/
├── build/               # Файлы для Docker
├── cmd/app/             # Основной пакет приложения
├── config/              # Файлы конфигурации
├── internal/app/        # Логика приложения
├── migrations/          # Файлы миграций БД
├── pkg/                 # Общие пакеты
├── .gitignore
├── docker-compose.yaml  # Конфигурация Docker Compose
├── go.mod               # Зависимости Go
├── go.sum               # Контрольные суммы зависимостей
```

## API Endpoints

### 1. Добавление пользователя
`POST /users`

**Тело запроса:**
```json
{
  "email": "ivan.petrov@example.com",
  "password": "password123",
}
```

### 2. Получение пользователя по ID
`GET /users?id={id}`

### 3. Обновление пользователя
`PUT /users`

**Тело запроса:**
```json
{
  "id": 1,
  "email": "ivan.petrov@example.com",
  "password": "new_password",
}
```

### 4. Удаление пользователя
`DELETE /users?id={id}`
