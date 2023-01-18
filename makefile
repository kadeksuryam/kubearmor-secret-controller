BINARY_NAME=kubearmor-secret-controller

all: build run

build: clean
	go fmt *.go
	go mod tidy
	go build -o ${BINARY_NAME}

clean:
	go clean

run:
	./${BINARY_NAME}
