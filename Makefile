run:
	go run ./cmd/main.go
build:
	go build -o hackers-list ./cmd/main.go
compose-build:
	docker-compose build
compose-up:
	docker-compose up