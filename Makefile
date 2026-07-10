USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)
MIGRATION_PATH=./db/migrations
DB_URL=postgres://postgres:pass@localhost:5432/my_app_db?sslmode=disable
GO=docker exec -it gin-api go

.PHONY: help
	
help:
	@echo "Available commands:"
	@echo "  migration-create name=<migration_name>     : Create a new migration"
	@echo "  migration-up                               : Apply pending migrations"
	@echo "  migration-down                             : Rollback the last migration"
	@echo "  migration-force V=<version>                : Force a specific migration version"
	@echo "  migration-drop                             : Drop all migrations, clean database"
	@echo "  generate-api                               : Generate API code from GraphQL schema"
	@echo "  generate-query                             : Generate GORM query code"

migration-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) $(name)

migration-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migration-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down

migration-force:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" force $(V)

migration-drop:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" drop

generate-api:
	$(GO) run github.com/99designs/gqlgen generate

generate-query:
	gorm gen -i ./internal/model/ -o ./internal/query/
