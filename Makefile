protoc:
	protoc --go_out=pd --go-grpc_out=pd \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=proto proto/*.proto \
		--grpc-gateway_opt=paths=source_relative --grpc-gateway_out=pd 	\
		--openapiv2_out=documents/swagger --openapiv2_opt=allow_merge=true,merge_file_name=myswagger

server:
	go run main.go

build:
	go build -o bin/server main.go

.PHONY: protoc server