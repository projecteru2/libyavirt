TEST := go test -count=1 -race -cover

PKGS := $$(go list ./... | grep -v vendor/)

.PHONY: all test build grpc

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
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

grpc:
	protoc --go_out=. \
			--go-grpc_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
			grpc/gen/yavirtd.proto

deps:
	go mod tidy

test:
ifdef RUN
	$(TEST) -v -run='${RUN}' $(PKGS)
else
	$(TEST) $(PKGS)
endif
