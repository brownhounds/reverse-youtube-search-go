init:
	cp .env.example .env

run:
	go run main.go

run_docker:
	docker compose up -d --build --force-recreate

build:
	GOOS=linux go build -ldflags="-s -w" -o ./bin/api main.go

lint:
	golangci-lint run --fast

fix:
	golangci-lint run --fix
