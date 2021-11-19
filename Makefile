build:
	go build -o ./bin/crypt main.go

test:
	go test -cover ./...
