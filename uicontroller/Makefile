GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=uicontroller

build: 
	$(GOBUILD) -v ./...

install:
	$(GOINSTALL) -v ./...
	
run:
	$(GORUN) cmd/uicontroller/main.go

