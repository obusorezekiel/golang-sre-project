# Makefile for a CRUD API built with Go, Gin, and GORM

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GORUN=$(GOCMD) run
GOMOD=$(GOCMD) mod
MAIN_NAME=main.go
BINARY_NAME=gin-gorm-crud

# Default task
all: build run

# Run go mod tidy to ensure dependencies are correct
deps:
	$(GOMOD) tidy

build: deps
	$(GOBUILD) -o $(BINARY_NAME) -v .

run:
	./$(BINARY_NAME)
