include .env

.PHONY: run
run:
	go run ./cmd

.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DB_DSN} up

.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests'
	go test -race -vet=off ./...

.PHONY: build
build:
	@echo 'Building cmd/'
	go build -ldflags='-s' -o=./bin/server ./cmd
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/server ./cmd


.PHONY: build/alwaysdata
build/alwaysdata:
	@echo 'Building cmd/ for hosting platform alwaysdata'
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s' -o=./bin/alwaysdata/server ./cmd
