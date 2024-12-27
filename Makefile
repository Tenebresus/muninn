.PHONY: all build clean run stop

all: build

restart: clean build run

build:
	@mkdir ./bin
	@go build -o ./bin/muninn_apiserver ./cmd/apiserver
	@go build -o ./bin/muninn_agent ./cmd/agent
	@go build -o ./bin/muninn_incubator ./cmd/incubator
	@go build -o ./bin/muninn_scavenger ./cmd/scavenger
	@echo "Succesfully built binaries"

clean:
	@rm -rf ./bin
	@echo "Cleaned up ./bin"

run:
	@./scripts/start.sh

stop:
	@./scripts/stop.sh
