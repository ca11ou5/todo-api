## Run tests
```bash
go test -cover -v ./...
```
or if you have make utility
```bash
make runtests
```

## Run app
There may be problems accessing the Docker registry, i use VPN
```bash
docker compose -f ./deployments/docker-compose.yml -p todo-app up --no-deps --force-recreate --build
```
or if you have make utility
```bash
make compose
```

## Swagger docs
http://localhost:8080/swagger/index.html


