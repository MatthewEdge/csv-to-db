# Tabload

## Run Locally

The application exposes the `/load` HTTP URI which supports Multipart uploads of the desired CSV file:

```
./knock --port 8080 &
curl -F upload=@file.csv http://localhost:8080/load
```

## Deployment

The application can be packaged in the provided Dockerfile and deployed to any Docker orchestrator:

```
docker build -t knock-app .
docker run --rm -p 8080:8080 knock-app:latest
```
