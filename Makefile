GOOS := darwin
GOARCH := arm64

.PHONY: build
build:
	@echo "Building application..."
	go build -o build/app main.go
