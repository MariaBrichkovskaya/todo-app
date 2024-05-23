FROM golang:alpine

COPY ./ ./
RUN go mod download

# Строим приложение
RUN go build -o todo-app ./cmd/main.go

# Копируем файл конфигурации
#COPY configs/config.yaml configs/

# Определяем порт для приложения
EXPOSE 9090

# Запускаем приложение
CMD ["./todo-app"]
