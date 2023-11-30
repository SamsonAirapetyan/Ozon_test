.PHONY: protos

protos:
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=import protos/link.proto

psql:
	echo STORAGE_TYPE=psql>.env
	go run cmd/app/main.go
in-memory:
	echo STORAGE_TYPE=inMemo>.env
	go run cmd/app/main.go
tests:
