all: client server

protoc:
	@echo "Generating Go files"
	cd proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

server: protoc
	@echo "Building server"
	go build -o bin/server cmd/main.go

client: protoc
	@echo "Building client"
	go build -o bin/client client_cmd/main.go

run_server: server
	@echo "starting server..."
	go run cmd/main.go

run_client: client
	@echo "executing client..."
	go run client_cmd/main.go

clean:
	go clean github.com/Connect2naga/go-examples/grpc_unary/...
	rm -rf bin

.PHONY: client server protoc