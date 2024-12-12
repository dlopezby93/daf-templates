BINARY_NAME=bootstrap

build:
	(cd cmd && GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./bin/bootstrap main.go)

wire:
	(cd internal/di && go run github.com/google/wire/cmd/wire)
