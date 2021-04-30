FROM golang:latest

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8000

CMD ["./go-fib"]