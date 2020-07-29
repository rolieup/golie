GO=GO111MODULE=on go

all: build test

build: golie golied

golie:
	$(GO) build -v cmd/golie/main.go

golied:
	$(GO) build -v cmd/golied/main.go

test:
	$(GO) test ./...

.PHONY: vendor
vendor:
	$(GO) mod tidy
	$(GO) mod vendor
	$(GO) mod verify
