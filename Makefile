USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)
MIGRATION_PATH=./db/migrations
DB_URL=postgres://postgres:pass@localhost:5432/my_app_db?sslmode=disable
GO=docker exec -it gin-api go

migration-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) $(name)

migration-up:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" up

migration-down:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" down

generate-api:
	$(GO) run github.com/99designs/gqlgen generate

generate-query:
	gorm gen -i ./internal/model/ -o ./internal/query/
