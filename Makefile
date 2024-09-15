dev:
	@if ! command -v air &> /dev/null; then \
    	echo "Please install air (go install github.com/air-verse/air@latest)"; \
    	exit 1; \
	fi
	air

run: compile
	./bin/main

compile:
	mkdir -p bin
	@templ generate
	go build -ldflags "-w" -o bin/main cmd/web/main.go

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

seed:
	go run cmd/seed/main.go

test:
	go test ./... -race -count 1 -timeout 600s | grep -v 'no test files'

generate-mock:
	@if ! command -v mockery &> /dev/null; then \
    	echo "Please install https://vektra.github.io/mockery/latest/installation/#go-install"; \
    	exit 1; \
	fi

	mockery