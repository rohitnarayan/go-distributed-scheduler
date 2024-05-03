.PHONY: clean build

# Variables
BINARY_NAME=out/go-distributed-scheduler

# Build the binary
build:
	go build -o $(BINARY_NAME)

# Clean up build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)
