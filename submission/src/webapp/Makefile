build:
	GOOS=linux GOARCH=amd64 go build -o bin/webapp main.go
	docker build . -t spca_webapp:latest

test:
	go test -v ./...