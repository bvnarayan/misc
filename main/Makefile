GO_FILES     := $(shell find . -name "*.go")
OS           := darwin
ARCH         := amd64

build:
	go vet ./...
	golint ./...
	GOOS=$(OS) GOARCH=$(ARCH) go build battleship.go 
	#go test

$(GOPATH)/bin/golint:
	go get -v github.com/golang/lint/golint
