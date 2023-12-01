.PHONY: protos

protos:
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=require_unimplemented_servers=false:. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --go-grpc_opt=paths=import -I protos/google/api -I protos/shortLink protos/shortLink/link.proto
psql:
	echo STORAGE_TYPE=psql>.env
	docker compose --profile db up --build
	#go run cmd/app/main.go
in-memory:
	echo STORAGE_TYPE=inMemory>.env
	docker compose --profile memory up --build
	#go run cmd/app/main.go
tests:
	go test -cover -v ./internal/service
