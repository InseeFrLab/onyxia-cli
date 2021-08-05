.PHONY: help # This help message
help:
	@grep '^.PHONY: .* #' Makefile \
	| sed 's/\.PHONY: \(.*\) # \(.*\)/\1\t\2/' \
	| expand -t20 \
	| sort

.PHONY: dev-prepare # Install go tools and pre-commit hooks
dev-prepare:
	go get github.com/securego/gosec/v2/cmd/gosec
	go get github.com/mgechev/revive
	go get golang.org/x/lint/golint
	#go get github.com/go-critic/go-critic/cmd/gocritic
	#go get github.com/akrennmair/go-imports
	#go get github.com/golangci/golangci-lint/cmd/golangci-lint

	pre-commit install

.PHONY: pre-commit # Run pre-commit compliance tests
pre-commit:
	pre-commit run --all-files

.PHONY: test # Run go test
test:
	go test

onyxiactl: test
	go build -o onyxyactl main.go

.PHONY: all # lint, test and build
all: onyxiactl
	@echo
