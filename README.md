# Tabload

## Run Locally

The app currently supports `PostgreSQL` for loading. A Docker image is a handy way to test:

`docker run --name local-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres`

Then, to load a CSV file:

`./knock-app -path "./test.csv" -host "localhost" -port "5432" -user "postgres" -pass "mysecretpassword"`

You can also run the application as a HTTP API which will allow a file upload interface as well. The application 
exposes the `/load` HTTP URI for CSV uploads

```
./knock-app -server -host "localhost" -port "5432" -user "postgres" -pass "mysecretpassword"
curl -F upload=@test.csv http://localhost:8080/load
```

## Deployment

The application can be packaged in the provided Dockerfile and deployed to any Docker orchestrator:

```
docker build -t knock-app .
docker run --rm -p 8080:8080 knock-app:latest
```
