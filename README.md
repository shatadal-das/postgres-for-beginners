# Start a Postgres database using Docker:

*Start docker desktop before running the following commands.*

## Command formats:

```bash
docker network create <network-name>
```

```bash
docker run --name <container_name> -e POSTGRES_USER=<username> -e POSTGRES_PASSWORD=<password> -d -p 5433:5432 --network=<network-name> postgres
```
URI for connecting to the database is `postgres://<username>:<password>@localhost:5433/<database_name>?sslmode=disable`

## For example:

```bash
docker network create postgres-network
```

```bash
docker run --name postgres-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5433:5432 --network=postgres-network postgres
```

Now, you can connect to the database using the following information:
- Host: `localhost`
- Port: `5433`
- Username: `postgres`
- Password: `postgres`
- Database: `postgres`

Your URI is `postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable`

## Files:
- `.env`: The environment file that stores confidential information (like the database URI).
- `main.go`: The main file that runs the application.
- `services/sql.go`: The SQL service that handles SQL operations.
- `sql/user/createTable.sql`: The SQL file that creates the user table.
- `sql/user/insertData.sql`: The SQL file that inserts data into the user table.
- `sql/user/findOne.sql`: The SQL file that finds one user from the user table.