build:
	go build ./cmd/lipsync

install:
	go install ./cmd/lipsync

fmt:
	go mod tidy
	go fmt ./...

check:
	golangci-lint run ./...
