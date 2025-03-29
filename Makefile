export ENV=development

server:
	go run cmd/*

run-compose:
	docker compose -f deployment/docker/docker-compose.yaml up -d

build-docker:
	docker build -t hicon . -f ./deployment/docker/Dockerfile -D

update-submodules:
	git submodule update --init --recursive && \
	git submodule foreach git checkout $(branch) && \
	git submodule foreach git pull origin $(branch)

proto:
	protoc --go_out=./hicon-sm --go-grpc_out=./hicon-sm ./hicon-sm/*.proto
