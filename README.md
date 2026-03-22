# Структура проекта route-nn
  ├── cmd/
  │   └── route-nn/
  │       └── main.go       # Точка входа (собирает зависимости и запускает app.Run)
  ├── config/               
  │   └── config.yaml       # Файл с настройками (URL 1C, Порты, Тайм-ауты)
  ├── internal/
  │   ├── app/              # Оркестратор (Run). Связь бизнес-логики и транспорта
  │   ├── client/           # Клиент для соединения с провайдером (1С)
  │   │   └── http_client.go
  │   ├── config/           # Пакет для чтения и парсинга config.yaml в структуру Go
- │   ├── delivery/         # Primary adapters (ведущие адаптеры). Слой доставки.
 #│   │   └── html/         # Пакет SSR (Handler + Templates)
  │   │       ├── handler.go
  │   │       └── templates/
  │   │            └── index.html
  │   ├── domain/           # Бизнес-сущности (Entity), интерфейсы и логика расчетов
- │   ├── repository/       # Secondary adapters (ведомые адаптеры). "Выход": 1С, БД и т.д.
 #│   │   └── provider/     # Конкретная реализация для 1С
  │   │       └── client.go
  │   └── usecase/          # Сценарии использования (UseCase)
