FROM golang:1.14.1-alpine3.11 AS builder
WORKDIR /app
COPY . /app
RUN go build -o app ./cmd/app

FROM alpine:3.11
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8081
ENTRYPOINT ["./app"]
