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


.PHONY: new_migration migrate_up migrate_down
