.PHONY: install tests build-js install-js

install-js:
	npm install
build-js:
	npm run build

install:
	go mod tidy && go mod vendor
	npm install && npm run build

tests:
	go test -covermode=set ./... -coverprofile=coverage.txt