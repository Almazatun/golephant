FROM golang:1.18.0-alpine3.14

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY .env /app

COPY Makefile /app

RUN apk update && apk add make

RUN make install

COPY . .

RUN go build -o main ./cmd

WORKDIR /dist

RUN cp /app/main .

EXPOSE 3000

CMD ["/dist/main"]