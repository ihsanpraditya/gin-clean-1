# Gin Clean 01

Boilerplate yang masih belum rapih untuk Gin [GraphQL](https://graphql.org/) dengan Docker, menggunakan:  
1. [GORM](https://gorm.io/)
2. [CLI nya](https://gorm.io/cli)
3. [golang-migrate](https://github.com/golang-migrate/migrate)
4. [PostgreSQL](https://www.postgresql.org/)
5. Layered Architecture (handler, service, repository)

## Quick Start

Just run `docker compose up -d` with a good internet connection.

## Generating JWT Key

````bash
openssl rand -base64 32
````
