BINARY = poultracker
VERSION := $(shell git describe --always --long --dirty)

default: build

build:
	CGO_ENABLED=1 go build -o ${BINARY}-${VERSION}-$(shell go env GOOS)-$(shell go env GOARCH) ./cmd/poultracker

linux: 
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ${BINARY}-${VERSION}-linux-amd64 ./cmd/poultracker

darwin:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}-${VERSION}-darwin-amd64 ./cmd/poultracker
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ${BINARY}-${VERSION}-darwin-darwin ./cmd/poultracker

windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ${BINARY}-${VERSION}-windows-amd64 ./cmd/poultracker

tailwind:
	tailwindcss -c ./web/tailwind.config.js -o ./web/static/css/tailwind.css --minify

tailwind-watch:
	tailwindcss -c ./web/tailwind.config.js -o ./web/static/css/tailwind.css --watch

run:
	@go run ./cmd/poultracker

live-backend-reload:
	@air

lint:
	@golangci-lint run

vulncheck:
	@govulncheck ./...

clean:
	go clean
	rm -f ${BINARY}-${VERSION}-$(shell go env GOOS)-$(shell go env GOARCH)
	rm -f ${BINARY}-${VERSION}-linux-amd64
	rm -f ${BINARY}-${VERSION}-darwin-amd64
	rm -f ${BINARY}-${VERSION}-darwin-darwin
	rm -f ${BINARY}-${VERSION}-windows-amd64

.PHONY: build tailwind tailwind-watch run live-backend-reload lint vulncheck
