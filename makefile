.PHONY: lint
.PHONY: install-tools

lint:
	golangci-lint run

install-tools:
	go install github.com/matryer/moq@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest