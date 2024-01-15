FROM golang:1.21-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY bin/start.sh .
COPY bin/wait-for .
COPY db/migration ./migration
RUN chmod +x /app/start.sh /app/wait-for

EXPOSE 8080
ENTRYPOINT [ "/app/start.sh" ]
CMD [ "/app/main" ]