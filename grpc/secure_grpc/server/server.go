// Package grpc_example contains ...
package server

import (
	"context"
	chat "github.com/connect2naga/go-examples/grpc/secure_grpc/proto"
	"log"

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

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func (s *Server) SayHello(ctx context.Context, msg *chat.MessageReq) (*chat.MessageRes, error) {
	log.Printf("got the message from client Msg:%s", msg.Body)
	return &chat.MessageRes{Body: "Msg received"}, nil
}
