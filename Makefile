# Go parameters
GO := go
BINARY_NAME := pomo
VERSION := 1.0.0

# Build flags
LDFLAGS := -ldflags="-w -s -X main.version=$(VERSION)"

.PHONY: all build install clean

all: clean install

build:
	@echo "Building $(BINARY_NAME) version $(VERSION)"
	CGO_ENABLED=0 $(GO) build --trimpath $(LDFLAGS) -o $(BINARY_NAME) ./main.go

install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin/$(BINARY_NAME)"
	sudo mv ./$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

clean:
	@echo "Cleaning up"
	rm -f $(BINARY_NAME) # Remove from working directory
	sudo rm -f /usr/local/bin/$(BINARY_NAME) # Remove from /usr/local/bin
