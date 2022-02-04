# grpc-example : Server Streaming 
in server side streaming client has to listen for the response  and server will send the response continuously and end send EOF

## Protobuf
in protobuf, we have to create the service rpc response as stream, so that server will open the stream until the EOF
```protobuf
service ChatSteamService{
  rpc SayHello(MessageReq) returns(stream MessageRes){}
}
```

## Server
in server, client request should read and send response in stream, here we used goroutines to send msg to clients.
```go
func (ss StreamServer) SayHello(req *msg_stream.MessageReq, svc msg_stream.ChatSteamService_SayHelloServer) error {
	log.Printf("fetch response for : %s", req.Body)

	//use wait group to allow process to be concurrent
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(count int64) {
			defer wg.Done()
			//time sleep to simulate server process time
			time.Sleep(time.Duration(count) * time.Second)
			resp := msg_stream.MessageRes{Body: fmt.Sprintf("response #%d For request:%s", count, req.GetBody())}
			if err := svc.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
			log.Printf("finishing request number : %d", count)
		}(int64(i))
	}

	wg.Wait()
	return nil
}

```

##Client
client should listen for stream message
```go
responseStream, err := c.SayHello(context.TODO(), &msg_stream.MessageReq{Body: "Nagarjuna"})
	if err != nil {
		log.Fatalf("unable to pocess request to server , error: %v", err)
	}

	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		for {
			// we can also create blocking recv call until sever send.
			data, err := responseStream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("failed while recieving message stream, %v", err)
			}

			log.Printf("data in stream : %s", data.Body)
		}
	}()

```


# execution 
generate server binary, 50052 is the port 
``` make 
 make server
 ./bin/server
--------------------------------------------logs ------------------------------------
 snagarju@snagarju  ~/.../go-examples/grpc_server_stream   grpc-serverStream ●  make server 
Generating Go files
cd proto && protoc --go_out=. --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto
Building server
go build -o bin/server server/server.go
 snagarju@snagarju  ~/.../go-examples/grpc_server_stream   grpc-serverStream ●  ./bin/server 
2022/02/04 12:08:35 fetch response for : Nagarjuna
2022/02/04 12:08:35 finishing request number : 0
2022/02/04 12:08:36 finishing request number : 1
2022/02/04 12:08:37 finishing request number : 2
^C
 ✘ snagarju@snagarju  ~/.../go-examples/grpc_server_stream   grpc-serverStream ●  

```

generate client and execute the binary
```
make client
./bin/client

---------------------------------------------------logs----------------------------------
 snagarju@snagarju  ~/.../go-examples/grpc_server_stream   grpc-serverStream ●  make client 
Generating Go files
cd proto && protoc --go_out=. --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto
Building client
go build -o bin/client client/client.go
 snagarju@snagarju  ~/.../go-examples/grpc_server_stream   grpc-serverStream ●  ./bin/client 
2022/02/04 12:08:35 data in stream : response #0 For request:Nagarjuna
2022/02/04 12:08:36 data in stream : response #1 For request:Nagarjuna
2022/02/04 12:08:37 data in stream : response #2 For request:Nagarjuna
2022/02/04 12:08:37 all the messages recieved from server....


```




