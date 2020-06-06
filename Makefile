all: test build

install:
	go install

test:
	go test ./...

build:
	go build -v -ldflags="-w -s" -o output/app
