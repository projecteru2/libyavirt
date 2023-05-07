OS := $(shell uname)
ARCH := $(shell uname -m)

ifeq "$(ARCH)" "aarch64"
PROTOC_ARCH := aarch_64
else ifeq "$(ARCH)" "arm64"
PROTOC_ARCH := aarch_64
else
PROTOC_ARCH := $(ARCH)
endif

# We only support Linux and OSX now.
ifeq "$(OS)" "Linux"
PROTOC_OS := linux
else ifeq "$(OS)" "Darwin"
PROTOC_OS := osx
endif

PROTOC_VER := 21.7
PROTOC_REL := https://github.com/protocolbuffers/protobuf/releases
PROTOC_BASENAME := protoc-$(PROTOC_VER)-$(PROTOC_OS)-$(PROTOC_ARCH)
PROTOC_TARGET_DIR := ${HOME}/.local/$(PROTOC_BASENAME)

TEST := go test -count=1 -race -cover

PKGS := $$(go list ./... | grep -v vendor/)

.PHONY: all test build setup setup-protoc grpc

lint: format
	PATH=${HOME}/go/bin:${PATH} \
	golangci-lint run --skip-dirs-use-default --skip-dirs=thirdparty

format:
	go vet $(PKGS)
	go fmt $(PKGS)

setup: setup-protoc
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/vektra/mockery/v2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

setup-protoc: FILENAME = $(PROTOC_BASENAME).zip
setup-protoc: DOWNLOAD_URL = $(PROTOC_REL)/download/v$(PROTOC_VER)/$(FILENAME)
setup-protoc: TARGET_DIR = $(PROTOC_TARGET_DIR)
setup-protoc:
	curl -o /tmp/${FILENAME} -L ${DOWNLOAD_URL}
	rm -fr ${TARGET_DIR}.new
	mkdir -p ${TARGET_DIR}
	unzip -o /tmp/${FILENAME} -d ${TARGET_DIR}.new
	if [ -e ${TARGET_DIR} ]; then rm -fr ${TARGET_DIR}.orig && mv ${TARGET_DIR} ${TARGET_DIR}.orig; fi
	mv ${TARGET_DIR}.new ${TARGET_DIR}

grpc:
	PATH=${HOME}/go/bin:${PROTOC_TARGET_DIR}/bin:${PATH} \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/gen/yavirtd.proto

deps:
	go mod tidy

test:
ifdef RUN
	$(TEST) -v -run='${RUN}' $(PKGS)
else
	$(TEST) $(PKGS)
endif
