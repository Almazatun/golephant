FROM golang:1.18.0-alpine3.14

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY .env /app

RUN go mod download

COPY . .

RUN go build -o main ./cmd

WORKDIR /dist

RUN cp /app/main .

EXPOSE 3000

CMD ["/dist/main"]