build:
	go build ./cmd/lipsync

install:
	go build ./cmd/lipsync

fmt:
	go mod tidy
	go fmt ./...
