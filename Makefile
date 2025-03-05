export ENV=development

server:
	go run cmd/main.go
run-services:
	docker compose -f deployment/docker/docker-compose.yaml up -d
