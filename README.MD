# Guide

first start this project

```bash
cp .env.example .env
```

## Install mod

```bash
go mod tidy
```

### Docker

```bash
docker build -t go-gin
```

```bash
docker-compose up -d
```

```bash
docker run -p 6476:6476 --name go-gin-app go-gin-app
```

```bash
docker-compose up -d
```

#### RUN

```bash
go run .
```

##### TEST

```bash
go test ./test
```
