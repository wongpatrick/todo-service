FROM golang:1.21

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o app main.go

EXPOSE 8080:8080

CMD ["./app"]