# RULES to CONTRIBUTE

## Structure

- **.env.example**: an example environment configuration file that can be used as a template for creating a .env file.
- **gqlgen.yml**: configuration file for gqlgen, a GraphQL server library for Go.
- **model.conf**: casbin model configuration file.
- **policy.csv**: casbin policy file that defines access control rules.
- **.air.toml**: configuration file for Air, a live-reloading tool for Go applications.
- **Makefile**: a file containing build commands and scripts for the application.
- **internal/**: contains the internal application code, including models, queries, and business logic.
  - **internal/database/db.go**: contains database connection and initialization code.
  - **internal/model/**: contains the data models for the application.
  - **internal/query/**: contains the generated query code for interacting with the database.
  - **internal/service/**: contains the business logic and service layer of the application.
  - **internal/handler/**: contains the HTTP handlers for the application.
  - **internal/config/**: contains configuration files and utilities for the application.
  - **internal/util/**: contains utility functions and helpers used throughout the application.
  - **internal/router/router.go**: contains the routing configuration for the application.
  - **internal/middleware/**: contains middleware functions for the application.
  - **internal/validator/**: contains validation logic for the application.
- **cmd/**: contains the main application entry point.
- **docs/**: contains documentation files, including this CONTRIBUTING.md file.
- **graphql/**: contains GraphQL schema files and related code.
  - **graphql/schema/\*.graphqls**.graphqls: contains the GraphQL schema definitions.
  - **graphql/resolver/\*.resolvers.go**: contains the resolver functions for the GraphQL schema.
  - **graphql/generated.go**
  - **graphql/resolver.go**
  - **graphql/model/models_gen.go**
- **db/migrations/**: contains database migration files for managing schema changes.
- **pkg/**: contains reusable packages and libraries that can be used across different parts of the application.
- **scripts/**: contains scripts for various tasks, such as database migrations or code generation.
- **test/**: contains test files and test-related utilities.

## Adding/Editing GraphQL Schema

1. Add/edit schema in `graphql/schema/*.graphqls`.
2. Run: 
```
go run github.com/99designs/gqlgen generate
```
3. Add or edit resolver in equivalent `graphql/resolver/*.resolvers.go`.
4. Test the GraphQL API using GraphQL Playground or any other GraphQL client.

## Adding/Editing Table

**1. Create a new migration file**
**2. Add new model in internal/model/ folder**
**3. Run [GORM CLI](https://gorm.io/cli/)**
```
gorm gen -i ./internal/model/ -o ./internal/query/
```
4. Define its graphql schema in `graph/`, follow instructions above.
5. Add/edit service in `internal/service/` folder.
6. Add/edit repository in `internal/repository/` folder.
