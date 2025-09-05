# ---------- BUILD STAGE ----------
FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12
RUN swag init -g main.go
RUN go build -o main main.go

# ---------- RUNTIME STAGE ----------
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/main .
COPY --from=builder /app/configs/docker/configs.json ./configs/
COPY --from=builder /app/.env .
COPY --from=builder /app/docs ./docs/

EXPOSE 9090
CMD ["./main"]
