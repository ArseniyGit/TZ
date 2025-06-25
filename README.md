Сервис для управления задачами с асинхронной обработкой, реализованный с использованием чистой архитектуры Роберта Мартина и принципов SOLID.

## API Endpoints

- `POST /api/v1/tasks` - Создать задачу
- `GET /api/v1/tasks` - Получить все задачи
- `GET /api/v1/tasks/{id}` - Получить задачу по ID
- `DELETE /api/v1/tasks/{id}` - Удалить задачу

  ##Запуск
  go mod tidy
  go run cmd/api/main.go

Сервер будет доступен на http://localhost:8080

## Примеры использования
  Создание задачи:
  curl -X POST http://localhost:8080/api/v1/tasks
  
  Получение статуса задачи:
  curl http://localhost:8080/api/v1/tasks/{task-id}
