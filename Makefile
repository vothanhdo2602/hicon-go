export ENV=development

server:
	go run cmd/main.go

run-services:
	docker compose -f deployment/docker/docker-compose.yaml up -d

update-submodules:
	git submodule update --init --recursive && \
	git submodule foreach git checkout $(branch) && \
	git submodule foreach git pull origin $(branch)

proto:
	protoc --go_out=./hicon-sm --go-grpc_out=./hicon-sm ./hicon-sm/*.proto
