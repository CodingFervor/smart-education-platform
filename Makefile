.PHONY: build run test clean swagger docker

APP_NAME=education-platform
BUILD_DIR=./bin

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/

run:
	go run ./cmd/api/

test:
	go test -v -cover ./...

clean:
	rm -rf $(BUILD_DIR)

swagger:
	swag init -g cmd/api/main.go -o docs/swagger

docker:
	docker-compose up -d --build

docker-down:
	docker-compose down

migrate:
	psql -f sql/init.sql $(DATABASE_URL)

lint:
	golangci-lint run
