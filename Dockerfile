FROM golang:1.21.6-alpine3.19 as builder

WORKDIR /application

COPY . .

RUN go mod tidy

RUN go build -o /app main.go

FROM alpine:3.19.1

WORKDIR /application

COPY --from=builder /app /app

EXPOSE 8000

ENTRYPOINT ["/app"]