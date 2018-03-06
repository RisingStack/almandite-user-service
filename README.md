# Almandite user service

## Prerequisites
- Dependencies
  - Make sure you have [dep tool](https://golang.github.io/dep/docs/installation.html) installed
  - Install dependencies with `dep ensure`
- PostgreSQL - run `docker-compose up -d`
- Environment variables
  - You can use a [`.env` file](https://github.com/joho/godotenv) for setting the environment variables
  - PG_URL: url of the postgres db (`postgres://postgres:w7o4bvt8ncp0ksd@localhost:5432/almandite?sslmode=disable`)
  - DEBUG_SQL: enable/disable query execution logs (`true/false`)
