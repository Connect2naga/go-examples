// Package main contains ...
package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	msg_stream "github.com/connect2naga/go-examples/grpc_server_stream/proto"

	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/4/22 10:20 AM
Project : grpc-example
File : server.go
*/
type StreamServer struct {
	msg_stream.UnimplementedChatSteamServiceServer
}

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

func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("enable to listen the port, error:%v", err)
	}
	s := StreamServer{}

	grpcServer := grpc.NewServer()
	msg_stream.RegisterChatSteamServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to server the grpc request")
	}

}
