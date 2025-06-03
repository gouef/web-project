APP_NAME := web
PROFILE_FILE := cpu.pprof
CALLGRIND_FILE := output.callgrind
PROFILE_DURATION := 30

.PHONY: install tests build-js install-js install-go build profile-build profile convert open clean run

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
	go build -ldflags="-s -w" -o $(APP_NAME) .

all: build-dev

profile-build:
	go build -gcflags=all="-N -l" -o $(APP_NAME) .

run:
	./$(APP_NAME)

profile:
	@echo "Profiling $(APP_NAME) for $(PROFILE_DURATION) seconds..."
	go tool pprof -seconds=$(PROFILE_DURATION) http://localhost:6060/debug/pprof/profile > $(PROFILE_FILE)

convert:
	@echo "Converting profile to callgrind format..."
	go tool pprof --callgrind ./$(APP_NAME) $(PROFILE_FILE) > $(CALLGRIND_FILE)

open:
	kcachegrind $(CALLGRIND_FILE)

clean:
	rm -f $(APP_NAME) $(PROFILE_FILE) $(CALLGRIND_FILE)
