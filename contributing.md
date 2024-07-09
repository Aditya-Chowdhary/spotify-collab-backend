# Project spotify-collab

## Getting Started
1. Create a psql database, edit the appropriate values in .env.example and rename to .env
2. In makefile, under db/migrations/up and /down commands, edit the dsn of the db to match your own.
3. Install the [migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md) tool with the appropriate driver tag(postgresql). Current command: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
4. Run make db/migrations/up to run the migrations against your database. If you don't have make you can also copy paste the command
5. If on Linux/Mac edit the make build command from main.exe to main
6. Install air `go install github.com/air-verse/air@latest`
7. Run `air` to start the server with live reloading

## Changes to DB
1. To create a new migration run db/migrations/new name={name}
2. To make new queries add in the appropriate file in internal/database/queries. Check sqlc docs to see how to structure queries
3. Run sqlc generate if any changes made in migrations/queries

## Handlers & Services
1. As of now just add a function below routes following the same format and add the appropriate route to the server. 


(Ignore below for now)
## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```