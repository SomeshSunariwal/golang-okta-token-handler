exec:
	bin/golang-okta-token-handler

build:
	go build -o bin/

run:
	go run main.go