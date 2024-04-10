ARCH=amd64
BINARY=bin/server

up:
	docker-compose up

down:
	docker-compose down

mysql:
	docker-compose exec mysql bash -c "mysql -u root -p"

dep:
	cd src && go mod download

run:
	cd src && go run main.go

build:
	cd src && GOARCH=${ARCH} GOOS=linux go build -o ../${BINARY}

clean:
	cd src && go clean && rm -rf ../${BINARY}

lint:
	golangci-lint run --enable-all
