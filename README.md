# Single Sign-On service (Go + gRPC)

## Технологии

    - Go 1.22+
    - gRPC + Protocol Buffers
    - PostgreSQL (позже)
    - Docker / docker-compose (позже)
    - Makefile для сборки и запуска

## Структура проекта

    - `cmd/auth-service/` — входная точка auth-сервиса.
    - `internal/auth/` — внутренняя логика авторизации/аутентификации.
    - `proto/` — gRPC контракты (.proto файлы)

## Как запустить (пока только auth-service)

    `make run-auth`

