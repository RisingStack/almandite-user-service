# Almandite user service

## Prerequisites
- Run `docker-compose up -d` to setup PostgreSQL database
- Make sure you have [dep tool](https://golang.github.io/dep/docs/installation.html) installed
- Create a `.env` file with the following environment variables
```
PG_URL=postgres://postgres:w7o4bvt8ncp0ksd@localhost:5432/almandite?sslmode=disable
DEBUG_SQL=true
```

## Startup
1. Install dependencies with `make deps`
2. Build the project with `make install`
3. Run the project with `almandite-user-server` command