FROM golang:alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && \
    apk add --no-cache git ca-certificates && \
    adduser -D -g '' appuser

WORKDIR /usr/src/app/knock/

COPY go.mod go.sum /usr/src/app/knock/
RUN go get && go mod vendor

COPY . /usr/src/app/knock/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /usr/src/app/knock/knock

FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/src/app/knock/knock /go/bin/knock

# Use an unprivileged user and start the app
USER appuser

CMD /go/bin/knock -server -host "${DB_HOST}" -port "${DB_PORT}" -user "${DB_USER}" -pass "${DB_PASS}"
