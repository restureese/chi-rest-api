FROM golang:1.21.6-bullseye

WORKDIR /application

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/application/binary"]