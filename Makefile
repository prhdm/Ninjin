go build : farm

all :  clean dependencies build

clean:
	@echo "Cleaning..."
	-rm -rf $(BUNARY_NAME)
	-rm -rf $(BINARY_UNIX)
	-$(GOCLEAN) -i
	@echo "Done cleaning"

dependencies:
	$(GOMOD) download
	$(GOMOD) vendor

proto-dependencies:
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	$(GOGET) github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	$(GOGET) github.com/golang/protobuf/protoc-gen-go

build-proto-go:
	find ./proto/messages -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}
	find ./proto/services -name "*.proto" 2>/dev/null | xargs realpath | xargs -I {} protoc $(PROTOC_IMPORT_PATH) --grpc-gateway_out=logtostderr=true:$(BUILD_PROTO_DIRECTORY) --swagger_out=logtostderr=true:$(BUILD_PROTO_DIRECTORY) --go_out=plugins=grpc:$(BUILD_PROTO_DIRECTORY) {}

farm:
	@echo "Building Farm"
	$(GOBUILD) ./...
	$(GOBUILD) -o $(BINARY_NAME) -v

GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=GO111MODULE=on $(GOCMD) mod
GOGET=GO111MODULE=on $(GOCMD) "get -u"
GOCLEAN=$(GOCMD) clean
BUILD_PROTO_DIRECTORY=../
GOOGLE_APIS_DIR="$$(find $(GOPATH) -wholename "*github.com/grpc-ecosystem/grpc-gateway*/third_party/googleapis" 2>/dev/null | head -n 1)"
PROTOC_IMPORT_PATH=-I${GOOGLE_APIS_DIR} -I $$PWD/proto -I/usr/local/include
BINARY_NAME=farm
BINARY_UNIX=$(BINARY_NAME)_unix
