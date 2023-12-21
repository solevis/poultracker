default: build

build:
	@CGO_ENABLED=1 go build ./cmd/poultracker

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

.PHONY: build tailwind tailwind-watch run live-backend-reload lint vulncheck
