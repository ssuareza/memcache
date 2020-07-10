ifeq ($(shell uname -s), Darwin)
    shasum=shasum -a256
else
    shasum=sha256sum
endif

build_args=-ldflags "-X main.versionString=dev" ./cmd/mapi
files=$(shell find cmd -type f)

.PHONY: test

build: build-linux build-darwin

build-linux: build/mapi-linux-amd64
build/mapi-linux-amd64: ${files}
	GOARCH=amd64 GOOS=linux go build -o $@ $(build_args)

build-darwin: build/mapi-darwin-amd64
build/mapi-darwin-amd64: ${files}
	GOARCH=amd64 GOOS=darwin go build -o $@ $(build_args)

start-darwin:
	build/mapi-darwin-amd64

start-linux:
	build/mapi-linux-amd64

test:
	@go test -v *.go
	@go test -v cmd/mapi/*.go

clean:
	@rm -rf test/*
	@rm -rf build/*
