# grpc-example : 
Source : https://grpc.io/docs/languages/go/quickstart/

##Pre-req : 
Protobuf compilers: 
https://grpc.io/docs/protoc-installation/
``` RHEL
$ yum install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```
Below statement used to generate go code from protobuf for server and client.
`go get -u github.com/golang/protobuf/protoc-gen-go`





##ProtoBuff File
**Step 3** : Create package of the protobuf file
```protobuf
syntax = "proto3";
package chat;
option go_package = "./chat";
```

**Step 2** : Create request and response for the API
```protobuf
message MessageReq{
  string body = 1;
}

message MessageRes{
  string body = 1;
}

```
**Step 3** : Create the service & its rpc calls
```protobuf
service ChatService{
  rpc SayHello(MessageReq) returns(MessageRes){}
}

```

**Step 4** : generate the code for the protobuf file
```protobuf
    protoc --go_out=./chat --go_opt=paths=source_relative \
    --go-grpc_out=./chat --go-grpc_opt=paths=source_relative \
    chat.proto
```

##Server :

**Step 1** : Create the server with listen port
```go
	lis, err:= net.Listen("tcp",":50052")
	if err != nil{
		log.Fatalf("Unable to open port :50052 for listen, failed with error: %v", err)
	}
```

**Step 2**: Create Grpc Server 
```go
gServer := grpc.NewServer()
```

**Step 3** : Attach listen port to gRPC server to server the requests
```go
if err := gServer.Serve(lis); err != nil{
		log.Fatalf("Failed to start GRPC Server")
	}
```


# execution 
generate server binary, 50052 is the port 
``` make 
 make server
 ./bin/server
--------------------------------------------logs ------------------------------------
snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ●  make server
Generating Go files
cd proto && protoc --go_out=. --go-grpc_out=. \
--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto
Building server
go build -o bin/server server/server.go
snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ● 
snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ● 
snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ● 
snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ●  ./bin/server
2022/02/04 11:47:31 got the message from client Msg:Hi this is Nagarjuna....
```

generate client and execute the binary
```
make client
./bin/client

---------------------------------------------------logs----------------------------------
 ✘ snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ●  make client 
Generating Go files
cd proto && protoc --go_out=. --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto
Building client
go build -o bin/client client/client.go
 snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ●  ./
bin/    client/ proto/  server/ vendor/ 
 snagarju@snagarju  ~/.../go-examples/grpc/unary   grpc_simple_ex ●  ./bin/client 
2022/02/04 11:47:31 response from server : Msg received

```




