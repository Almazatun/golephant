FROM golang:1.17-alpine3.15

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY .env /app

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /app/main .

EXPOSE 3000

CMD ["/dist/main"]