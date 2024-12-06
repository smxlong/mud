.PHONY: check
check: gen fmt vet lint test

.PHONY: gen
gen:
	go generate ./...

.PHONY: fmt
fmt: gen
	go fmt ./...

.PHONY: vet
vet: gen
	go vet ./...

.PHONY: lint
lint: gen
	golangci-lint run

.PHONY: test
test: gen
	go test ./...
