FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/golang-backend cmd/main.go

FROM golang:alpine

COPY --from=builder /app/bin/golang-backend .

ENTRYPOINT ["/app/bin/golang-backend"]