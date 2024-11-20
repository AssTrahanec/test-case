## Требования

- Docker
- Docker Compose

## Установка и запуск
## Docker Compose:
### 1: Запуск базы данных
```sh
docker-compose up test-case-db
```
### 2: Выполнение миграций
```sh
docker-compose run migrate
```
### 3: Запуск бэкенда
```sh
docker-compose up backend
```