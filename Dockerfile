# ---------- BUILD STAGE ----------
FROM golang:1.23.6 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь код
COPY . .

# Устанавливаем swag
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Генерируем swagger-документацию
RUN swag init

# Собираем бинарник
RUN go build -o main main.go


# ---------- RUNTIME STAGE ----------
FROM ubuntu:latest

WORKDIR /app

# Устанавливаем зависимости
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Создаём директорию для конфигов
RUN mkdir -p configs

# Создаём директорию для uploads
RUN mkdir -p uploads

# Копируем бинарник
COPY --from=builder /app/main .

# Копируем конфиги
COPY --from=builder /app/configs/docker/configs.json ./configs/

# Копируем .env
COPY --from=builder /app/.env .

# Копируем docs
COPY --from=builder /app/docs ./docs/

EXPOSE 9090

CMD ["./main"]
