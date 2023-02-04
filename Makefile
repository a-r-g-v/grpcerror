GOBIN := $(PWD)/.bin
PATH := $(GOBIN):$(PATH)

.PHONY: tools
tools:
	GOBIN=$(GOBIN) go install github.com/incu6us/goimports-reviser/v3@v3.3.0

.PHONY: fmt
fmt:
	goimports-reviser -rm-unused -set-alias -project-name github.com/a-r-g-v/grpcerror -format ./...;

.PHONY: test
test:
	go test ./...

.PHONY: codegen
codegen:
	go run ./internal/codegen > helper.go
	make fmt
