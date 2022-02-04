// Package grpc_example contains ...
package main

import (
	"context"
	"log"
	"net"

	chat "github.com/connect2naga/go-examples/grpc/unary/proto"

	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/2/22 3:10 PM
Project : grpc-example
File : server.go
*/

type Server struct {
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, msg *chat.MessageReq) (*chat.MessageRes, error) {
	log.Printf("got the message from client Msg:%s", msg.Body)
	return &chat.MessageRes{Body: "Msg received"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to open port :50052 for listen, error: %v", err)
	}

	s := Server{}
	gServer := grpc.NewServer()
	chat.RegisterChatServiceServer(gServer, &s)

	if err := gServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serv GRPC Server over port 50052, %v", err)
	}
}
