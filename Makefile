MIGRATION_PATH=./db/migrations
DB_URL=postgres://postgres:pass@localhost:5432/my_app_db?sslmode=disable

migration-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) $(name)

migration-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migration-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down
