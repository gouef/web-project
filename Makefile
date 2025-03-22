.PHONY: install tests build-js install-js install-go build

install-js:
	npm install
build-js:
	npm run build

install-go:
	go mod tidy && go mod vendor

install:
	go mod tidy && go mod vendor
	npm install && npm run build

tests:
	go test -covermode=set ./... -coverprofile=coverage.txt

build:
	go build -o app -ldflags="-s -w" main.go