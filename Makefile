include .env
BINARY_NAME=app

build:
	go build -o .bin/${BINARY_NAME} main.go
run:
	go build -o .bin/${BINARY_NAME} main.go
	chmod +x .bin/${BINARY_NAME}
	./.bin/${BINARY_NAME}
clean:
	go clean
	rm .bin/${BINARY_NAME}    
