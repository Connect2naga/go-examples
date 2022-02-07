// Package server contains ...
package main

import (
	"fmt"
	"io"
	"log"
	"net"

	bi_stream "github.com/connect2naga/go-examples/grpc/bidirectional_stream/proto"
	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/4/22 1:10 PM
Project : go-examples
File : server.go
*/

type biStream struct {
	bi_stream.UnimplementedChatBiSteamServiceServer
}

func (biStream) SayHello(srv bi_stream.ChatBiSteamService_SayHelloServer) error {
	log.Printf("session started with client...")
	ctx := srv.Context()

	count := 1
	oldMsg := ""
	for {

		//Call termination : exit if context is done or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			//Call termination : return will close stream from server side
			log.Println("exit")
			return nil
		}

		//There is a chance to get duplicate message,
		if req.Body == oldMsg {
			continue
		}
		oldMsg = req.Body

		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		log.Printf("Request Message got from clinet : %s", req.Body)
		resp := bi_stream.MessageRes{
			Body: fmt.Sprintf("response : %d for Req:%s", count, req.Body),
		}
		count = count + 1

		// send it to stream
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new message:%s ", resp.GetBody())
	}
	return nil
}

func main() {

	list, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("unable to list to the port :50052, error: %v", err)
	}

	grpcServer := grpc.NewServer()
	bi_stream.RegisterChatBiSteamServiceServer(grpcServer, biStream{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("unable to server the grpc service")
	}

}
