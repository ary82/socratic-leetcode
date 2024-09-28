all: build

build:
	@echo "building Go binary..."
	@go build -o main cmd/app/main.go

run: build
	@echo "running..."
	@./main

clean:
	@echo "cleaning binary..."
	@rm -f main

.PHONY: all build run clean
