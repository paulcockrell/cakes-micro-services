GOPATH:=$(shell go env GOPATH)
.PHONY: build
build:
	CGO_ENABLED=0 go build -o cake-service *.go

.PHONY: docker
docker:
	docker build . -t cake-service:latest
