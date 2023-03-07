FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY . . 

RUN go get -d -v ./...

RUN go build -o api .

EXPOSE 6476

CMD ["./api"]