.PHONY: all test 
BINARY_NAME := thumbpicker
PKG := ./cmd/server/main.go

.PHONY: build
build:
	templ generate && \
	go build -o $(BINARY_NAME) $(PKG)

.PHONY: run
run:build
	./$(BINARY_NAME)
	
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux $(PKG)

.PHONY: help 
help:
	@echo "  build         Compile the application"
	@echo "  run           Generate templates and run the application"
	@echo "  tidy          Tidy go.mod file"
	@echo "  clean         Remove artifacts"
	@echo "  build-linux   Build for Linux (amd64)"

