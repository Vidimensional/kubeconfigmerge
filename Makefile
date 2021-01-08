.DEFAULT_GOAL := build

.PHONY: test
test:
	go test -count=1 -race -cover --tags=integration ./...

.PHONY: build
build:
	go build -o ./bin/kubeconfigmerge