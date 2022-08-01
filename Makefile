.PHONY: build-cli
build-cli:
	go build -v -o ./bin/command ./cmd/cli
	chmod +x ./bin/command
