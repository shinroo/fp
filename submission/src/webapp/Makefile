build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/webapp main.go
	docker build . -t spca_webapp:latest

test:
	go test -v ./...