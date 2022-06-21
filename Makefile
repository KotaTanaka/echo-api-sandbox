ARCH=amd64
BINARY=bin/server

dep:
	go mod download

run:
	go run main.go

build:
	GOARCH=${ARCH} GOOS=linux go build -o ${BINARY}

clean:
	go clean
	rm ${BINARY}

lint:
	golangci-lint run --enable-all
