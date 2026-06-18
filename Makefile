VERSION := 1.0.0
BINARY_NAME := pregit

build:
	@echo "Building version $(VERSION)..."
	go mod tidy
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME)

install: build
	@echo "Installing to /usr/local/bin..."
	sudo mv $(BINARY_NAME) /usr/local/bin/

clean:
	rm -f $(BINARY_NAME)