DB_CONN=postgres://root:secret@localhost:5432/taskmanager?sslmode=disable

new_migration:
	@echo "Creating new migration..."
	@read -p "Enter migration name: " name; \
		if [ -z "$$name" ]; then \
			echo "Migration name cannot be empty"; \
			exit 1; \
		fi; \
		migrate create -ext sql -dir ./db/migrations $$name; \
		echo "Migration created successfully"
migrate_up:
	@echo "Running migrations..."
	@migrate -path ./db/migrations -database "$(DB_CONN)" -verbose up
	@echo "Migrations ran successfully"

migrate_down:
	@echo "Rolling back migrations..."
	@migrate -path ./db/migrations -database "$(DB_CONN)" -verbose down
	@echo "Migrations rolled back successfully"

sqlc:
	@echo "Generating sqlc..."
	@sqlc -f db/sqlc.yaml generate
	@echo "sqlc generated successfully"

mock:
	@echo "Generating mock..."
	@mockgen --package dbmock --destination db/mock/store.go taskmanager/db/model Store
	@echo "mock generated successfully"

build:
	@echo "Building server..."
	@docker-compose -p taskmanager -f docker/docker-compose.yml build
	@echo "Server built successfully"

start:
	@echo "Starting server..."
	@docker-compose -p taskmanager -f docker/docker-compose.yml up -d
	@echo "Server started successfully"

stop:
	@echo "Stopping server..."
	@docker-compose -p taskmanager -f docker/docker-compose.yml down
	@echo "Server stopped successfully"
	
test: 
	@echo "Running tests..."
	@go test -cover ./...
	@echo "Tests ran successfully"

.PHONY: new_migration migrate_up migrate_down sqlc mock build start stop test
