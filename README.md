# Gin Clean 01

Boilerplate yang masih belum rapih untuk Gin [GraphQL](https://graphql.org/) dengan Docker, menggunakan:  
1. [GORM](https://gorm.io/)
2. [CLI nya](https://gorm.io/cli)
3. [golang-migrate](https://github.com/golang-migrate/migrate)
4. [PostgreSQL](https://www.postgresql.org/)
5. Layered Architecture (handler, service, repository)

## Quick Start

1. Set up variables in `.env` file following the `.env.example` file.
2. Run `docker compose up -d` with a good internet connection.

## Generating JWT Key

````bash
openssl rand -base64 32
````

## Some GraphQL Queries

```graphql

{
  users {
    id
    name
    email
    roles {
      id
      name
    }
    isActive
  }
}

query {
  user(id: 1) {
    id
    name
    email
    roles {
      id
      name
    }
    isActive
  }
}

mutation {
    createUser(input: {
    name: "Ihsan",
    email: "ihsan@example.com"
    password: "islam123",
    confirmPassword: "islam123",
    roles: [1,2],
    isActive: true
  }) {
      id
      name
      email
  }
}

mutation {
  login(email: "ihsan@example.com", password: "islam123") {
    user {
      id
      name
      email
    }
    token
  }
}

mutation {
  updateUser(id: 1, input: {
    name: "hor",
    email: "hor@example.com",
    roles: [1,2],
    isActive: true
  }) {
    id
    name
    roles {
      id
    }
    isActive
  }
}

mutation {
  deleteUser(id: 1)
}
```