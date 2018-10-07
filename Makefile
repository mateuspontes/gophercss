.PHONY: test install

install:
	go get

test:
	go clean ./...
	go test ./...
