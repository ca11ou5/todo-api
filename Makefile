.PHONY: build
build:
	docker build -t todo-list -f ./build/Dockerfile ./

compose:
	docker compose -f ./deployments/docker-compose.yml -p todo-app up --no-deps --force-recreate --build

runtests:
	go test -cover -v ./...
