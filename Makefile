TEST := go test -count=1 -race -cover

PKGS := $$(go list ./... | grep -v vendor/)

.PHONY: all test build grpc

default: build

lint: format
	golint $(PKGS)

format:
	go vet $(PKGS)
	go fmt $(PKGS)

protoc:
	rm -fr /tmp/protoc-3.7.1-linux-x86_64.zip
	curl -Lv https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip -o /tmp/protoc-3.7.1-linux-x86_64.zip
	sudo unzip -o /tmp/protoc-3.7.1-linux-x86_64.zip -d /usr/local bin/protoc
	sudo unzip -o /tmp/protoc-3.7.1-linux-x86_64.zip -d /usr/local 'include/*'

grpc:
	protoc --go_out=plugins=grpc:. grpc/gen/yavirtd.proto

deps:
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor
	go get github.com/golang/protobuf/protoc-gen-go@v1.3

test:
ifdef RUN
	$(TEST) -v -run='${RUN}' $(PKGS)
else
	$(TEST) $(PKGS)
endif
