.PHONY: protos

protos:
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=import protos/link.proto