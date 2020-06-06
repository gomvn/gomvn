all: test build

install:
	go install

test:
	go test ./...

build:
	go build -v -ldflags="-w -s" -o output/app

docker-build:
	docker image build -t gomvn/gomvn .

docker-run: docker-build
	docker run -it -p 8080:8080 gomvn/gomvn
