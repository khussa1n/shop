FROM golang:1.22.1-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o shop ./cmd/main.go


FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/shop .
COPY config.yaml .
COPY migrations/pg/000001_init.up.sql /app/migrations/pg/000001_init.up.sql
COPY migrations/pg/000001_init.down.sql /app/migrations/pg/000001_init.down.sql
EXPOSE 8080
RUN apk update && apk add postgresql-client
CMD ["./shop 10,11,14,15"]