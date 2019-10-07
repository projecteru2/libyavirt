GOPATH := $(shell dirname $(shell dirname $(shell dirname $(shell dirname $(shell realpath .)))))
GO := GOPATH=$(GOPATH) go
TEST := $(GO) test -count=1 -race -cover

PKGS := $$($(GO) list ./... | grep -v vendor/)

.PHONY: all test build grpc

default: build

lint: format

format:
	$(GO) vet $(PKGS)
	$(GO) fmt $(PKGS)

grpc:
	protoc --go_out=plugins=grpc:. grpc/gen/yavirtd.proto

deps:
	GO111MODULE=on GOPATH=$(GOPATH) go mod download
	GO111MODULE=on GOPATH=$(GOPATH) go mod vendor

test:
ifdef RUN
	$(TEST) -v -run='${RUN}' $(PKGS)
else
	$(TEST) $(PKGS)
endif
