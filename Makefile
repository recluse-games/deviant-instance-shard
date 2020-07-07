build:
	go build -o ./bin/deviant-instance-shard ./cmd/deviant-instance-shard.go
run:
	go run ./cmd/deviant-instance-shard.go
test:
	go test ./... -cover