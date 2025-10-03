# Makefile to cross-compile Go project for Linux (static binary) from Windows

APP_NAME := book-management-backend
APP_NAME_SHORT := bk-manage
BUILD_DIR := build
GOOS := linux
GOARCH := amd64
CGO_ENABLED := 0

.PHONY: all clean build docker

all: build

build:
	@echo "Building $(APP_NAME) for $(GOOS)/$(GOARCH)..."
	@[ -d $(BUILD_DIR) ] || mkdir $(BUILD_DIR)
	set GOOS=$(GOOS) 
	set GOARCH=$(GOARCH) 
	set CGO_ENABLED=$(CGO_ENABLED) 
	go build -ldflags="-s -w -extldflags '-static'" -o $(BUILD_DIR)/$(APP_NAME) .

docker: build
	docker build -t $(APP_NAME_SHORT) .

clean:
	@echo "Cleaning build directory..."
	@[ -d ./build ] && rm -r $(BUILD_DIR)
