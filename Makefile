include .envrc

# Define variables for paths
APP_NAME = mekahtellyuh
SRC_DIR = ./cmd/web
MAIN_FILE = $(SRC_DIR)/main.go
MIGRATIONS_DIR = ./migrations
GO_FILES = $(shell find ./cmd ./internal -name '*.go')   # Exclude './ui' from Go files search
MIGRATION_FILES = $(shell find $(MIGRATIONS_DIR) -name '*.sql')

.PHONY: run/tests
run/tests: vet
	go test -v ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet: fmt
	go vet ./...

# Build and run the application
.PHONY: run
run: vet
	go run ./cmd/web -addr=${ADDRESS} -dsn=${FEEDBACK_DB_DSN}

# Database commands

## db/psql: Connect to PostgreSQL
.PHONY: db/psql
db/psql:
	psql ${FEEDBACK_DB_DSN}

## db/migrations/new name=$1: Create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: Apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${FEEDBACK_DB_DSN} up

## db/migrations/down-1: Undo the last migration
.PHONY: db/migrations/down-1
db/migrations/down-1:
	@echo 'Running down migrations...'
	migrate -path ./migrations -database ${FEEDBACK_DB_DSN} down 1

## db/migrations/fix: Fix a SQL migration if it's "dirty"
.PHONY: db/migrations/fix
db/migrations/fix:
	@echo 'Checking migration status...'
	@migrate -path ./migrations -database $${FEEDBACK_DB_DSN} version > /tmp/migrate_version 2>&1
	@cat /tmp/migrate_version
	@if grep -q "dirty" /tmp/migrate_version; then \
		version=$$(grep -o '[0-9]\+' /tmp/migrate_version | head -1); \
		echo "Found dirty migration at version $$version"; \
		echo "Forcing version $$version..."; \
		migrate -path ./migrations -database $${FEEDBACK_DB_DSN} force $$version; \
		echo "Running down migration..."; \
		migrate -path ./migrations -database $${FEEDBACK_DB_DSN} down 1; \
		echo "Running up migration..."; \
		migrate -path ./migrations -database $${FEEDBACK_DB_DSN} up; \
	else \
		echo "No dirty migration found"; \
	fi
	@rm -f /tmp/migrate_version

# Build the Go app
.PHONY: build
build: $(GO_FILES)
	@echo "Building the app..."
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(MAIN_FILE)

# Apply migrations
.PHONY: migrate
migrate:
	@echo "Applying database migrations..."
	# Loop through all migration files and apply them using `psql`
	$(foreach file,$(MIGRATION_FILES),psql -U youruser -d yourdb -f $(file);)

# Clean up the built binary
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Default target (build and run)
.PHONY: default
default: run
