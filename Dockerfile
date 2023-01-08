FROM golang:1.18.9-alpine3.17 AS builder

WORKDIR /app
COPY . .
RUN go build -o /app/main .

FROM alpine:3.17.0

WORKDIR /app
RUN apk --no-cache add ca-certificates musl-dev gcc
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]

# Path: users/docker-compose.yml