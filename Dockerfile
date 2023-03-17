FROM golang:1.19-alpine3.16 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -v ./test

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 6476

CMD ["./main"]