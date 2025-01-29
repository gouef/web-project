.PHONY: install tests build-js

build-js:
	npm install && npm run build

install:
	go mod tidy && go mod vendor
	npm install && npm run build

tests:
	go test -covermode=set ./... -coverprofile=coverage.txt