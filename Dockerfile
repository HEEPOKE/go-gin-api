FROM golang:1.19-alpine3.16 as builder

RUN apk add --no-cache git gcc ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.16

WORKDIR /root

COPY --from=builder /app/main .

EXPOSE 6476

CMD ["./main"]