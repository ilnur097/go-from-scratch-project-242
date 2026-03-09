```makefile
# Переменные
BINARY_NAME=myapp
MAIN_FILE=main.go

# .PHONY говорит Make, что это команды, а не названия файлов
.PHONY: all build run test clean deps

# Команда по умолчанию
all: deps build test

build:
@echo "🔨 Собираю бинарный файл..."
go build -o $(BINARY_NAME) $(MAIN_FILE)

run:
@echo "🚀 Запускаю проект..."
go run $(MAIN_FILE)

test:
@echo "🧪 Запускаю тесты..."
go test -v ./...

deps:
@echo "📦 Загружаю зависимости..."
go mod download

clean:
@echo "🧹 Очистка..."
go clean
rm -f $(BINARY_NAME)
