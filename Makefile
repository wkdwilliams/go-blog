dev:
	air

run: compile
	./bin/main

compile:
	mkdir -p bin
	@templ generate
	go build -o bin/main cmd/web/main.go

templ:
	source .env && templ generate --watch --proxy=http://localhost:$$PORT

migrate-create:
	@if ! command -v migrate &> /dev/null; then \
    	echo "Please install https://github.com/golang-migrate/migrate"; \
    	exit 1; \
	fi

	@read -p "Enter name of migration: " name; \
	migrate create -ext sql -dir migrations $$name

migrate-up:
	go run cmd/migrate/main.go -command up

migrate-down:
	go run cmd/migrate/main.go -command down