build:
	go build -o bin/golang-backend cmd/main.go

run: build
	bin/golang-backend

test:
	go test -v ./...