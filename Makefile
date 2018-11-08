.PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./hello-world/hello-world

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o hello-world/hello-world ./hello-world

run: build
	sam local start-api
