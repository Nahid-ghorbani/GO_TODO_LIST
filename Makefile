DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet:fmt
	go vet ./...

build:vet
	go build main.go

run:vet
	go run main.go
