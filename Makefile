
all: build

install:
		@go mod download

build:
		@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

run:
		@go run cmd/$(APP_NAME)/main.go