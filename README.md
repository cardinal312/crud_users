# Мини веб-приложение на Go

Это простое веб-приложение на языке Go, которое предоставляет несколько API-эндпоинтов для работы с данными.

Технологии:
- Go (Golang)
- Postgres (SQL db)
- Docker-compose

# Описание
Веб-приложение предоставляет API для выполнения следующих операций:

GET localhost:9090/users - Получить список всех пользователей.

GET localhost:9090/user{id} - Получить пользователя по ID.

POST localhost:9090/users - Создать нового пользователя.

PATCH localhost:9090/user{id} - Обновить информацию пользователя по ID.

DELETE localhost:9090/user{id} - Удалить пользователя по ID.

# Установка и запуск

Клонируйте этот репозиторий и запустите в терминале:

git clone https://github.com/cardinal312/crud_users.git

Запустите приложение в терминале с помошью Makefile:
--> make <<--

# Пример запроса JSON:

<img width="278" alt="Снимок экрана 2025-01-17 в 06 38 41" src="https://github.com/user-attachments/assets/ed189da3-0d68-4b9a-b2d3-eccbffe92279" />





