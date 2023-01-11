RPCOGNITO_BINARY=rpcognito

docker_start: build_rpcognito
	@echo "stops all running container"
	docker-compose down
	@echo "building(when required) and start docker images"
	docker-compose up --build
	@echo "docker images built and started"

build_rpcognito:
	@echo "building project"
	env GOOS=linux CGO_ENABLED=0 go build -o ${RPCOGNITO_BINARY} ./cmd
	@echo "done"