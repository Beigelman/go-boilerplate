SHELL=/bin/zsh
PROTO_SOURCES:=$(wildcard modules/**/interface/grpc/proto/*.proto)

proto:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	$(PROTO_SOURCES)
	 