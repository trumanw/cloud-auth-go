# This is a Makefile which maintains files automatically generated but to be
# shipped together with other files.

PKG=github.com/grpc-ecosystem/grpc-gateway
GOOGLEAPIS_DIR=third_party/googleapis
GO_SRC_PATH=$(GOPATH)/src
GOOGLEAPIS_DIR=third_party/googleapis
OPTIONS_PROTO=$(GOOGLEAPIS_DIR)/google/api/annotations.proto $(GOOGLEAPIS_DIR)/google/api/http.proto
PKGMAP=Mgoogle/api/annotations.proto=$(PKG)/$(GOOGLEAPIS_DIR)/google/api

# Add more protobuf src here so that they can be correctly compiled
PROTOC_SRC=pb/client_credentials.proto \
	pb/token.proto


SVCSRCS_GEN=protoc -I . \
	-I $(GO_SRC_PATH) \
	-I $(GO_SRC_PATH)/$(PKG)/$(GOOGLEAPIS_DIR) \
	--go_out=$(PKGMAP),plugins=grpc:. \
	$(PROTOC_SRC)

GWSRCS_GEN=protoc -I . \
	-I $(GO_SRC_PATH) \
	-I $(GO_SRC_PATH)/$(PKG)/$(GOOGLEAPIS_DIR) \
	--grpc-gateway_out=logtostderr=true:. \
	$(PROTOC_SRC)

protoc:
	$(SVCSRCS_GEN)
	$(GWSRCS_GEN)

build: protoc
	go build main.go -o build/oauth
