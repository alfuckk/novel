# Gin + Rel

Feature:

- Modular Project Structure.
- Full example including tests.
- Docker deployment.
- Compatible.

## Installation

### Running

1. Prepare `.env`.

    ```
    cp .env.sample .env
    ```

2. Start mysql and create database.

    ```
    docker-compose up -d
    ```

2. Prepare database schema.

    ```
    rel migrate
    ```

3. Build and Running

    ```
    make
    ```

## Project Structure

```
.
├── api
│   ├── handler
│   │   ├── todos.go
│   │   └── [other handler].go
│   └── middleware
│       └── [other middleware].go
├── bin
│   ├── api
│   └── [other executable]
├── cmd
│   ├── api
│   │   └── main.go
│   └── [other cmd]
│       └── main.go
├── db
│   ├── schema.sql
│   └── migrations
│       └── [migration file]
├── todos
│   ├── todo.go
│   ├── create.go
│   ├── update.go
│   ├── delete.go
│   ├── service.go
│   └── todostest
│       ├── todo.go
│       └── service.go
├── [other domain]
│   ├── [entity a].go
│   ├── [business logic].go
│   ├── [other domain]test
│   │   └── service.go
│   └── service.go
└── [other client]
    ├── [entity b].go
    ├── client.go
    └── [other client]test
        └── client.go
```
