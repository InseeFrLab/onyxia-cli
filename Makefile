.PHONY: help # This help message
help:
	@grep '^.PHONY: .* #' Makefile \
	| sed 's/\.PHONY: \(.*\) # \(.*\)/\1\t\2/' \
	| expand -t20 \
	| sort

.PHONY: pre-commit # Run pre-commit compliance tests
pre-commit:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
	pre-commit install
	pre-commit run --all-files

.PHONY: test # Run go test
test:
	go test

onyxiactl: test
	go build -o onyxyactl main.go

.PHONY: all # lint, test and build
all: pre-commit test onyxiactl
	@echo
