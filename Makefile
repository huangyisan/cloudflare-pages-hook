.PHONY: build-linux build-osx build-windows clean

VERSION=$(shell git rev-parse --short HEAD)
BUILD=$(shell date +%FT%T%z)

build-linux:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/linux/cloudflare-pages-hook

build-osx:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=darwin go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/darwin/cloudflare-pages-hook

build-windows:
	@GOARCH=amd64 CGO_ENABLED=1 GOOS=windows go build -ldflags "-s -w -X main.Version=${VERSION} -X main.Build=${BUILD}" -o bin/win/cloudflare-pages-hook

clean:
	@if [ -f bin/linux/cloudflare-pages-hook ] ; then rm -rf bin/linux/ ; fi
	@if [ -f bin/darwin/cloudflare-pages-hook ] ; then rm -rf bin/darwin/ ; fi
	@if [ -f bin/win/cloudflare-pages-hook ] ; then rm -rf bin/win/ ; fi
