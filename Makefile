build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ./bin/deviant-instance-shard ./cmd/deviant-instance-shard.go
	GOOS="windows" GOARCH="amd64" CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ./bin/deviant-instance-shard.exe ./cmd/deviant-instance-shard.go
run:
	go run ./cmd/deviant-instance-shard.go
test:
	go test ./... -cover
docker:
	sudo -E docker-compose up --build 