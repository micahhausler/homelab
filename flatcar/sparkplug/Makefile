

build:
	go build -o sparkplug .

test:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
