# Tabload

## Run Locally

The app currently supports `PostgreSQL` for loading. A Docker image is a handy way to spin up a database for testing:

`docker run --name local-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres`

Then, to load a CSV file:

`./tabload -path "./test.csv" -host "localhost" -port "5432" -user "postgres" -pass "mysecretpassword"`

You can also run the application as a HTTP API which will allow a file upload interface as well. The application 
exposes the `/load` HTTP URI for CSV uploads

```
./tabload -server -host "localhost" -port "5432" -user "postgres" -pass "mysecretpassword"
curl -F upload=@test.csv http://localhost:8080/load
```

## Docker

The provided Dockerfile, by default, spins the application up in HTTP mode with the database credentials passed through environment variables.
For convenience a `docker-compose.yml` file is provided to spin up the image and a Postgres image for testing. To use this:

```
docker-compose build
docker-compose up
```

## Deployment

The application can be packaged in the provided Dockerfile and deployed to any Docker orchestrator:

```
docker build -t tabload .
docker run --rm -p 8080:8080 tabload:latest
```
