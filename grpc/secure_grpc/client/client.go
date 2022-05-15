// Package client contains ...
package main

import (
	"context"
	"log"

	chat "github.com/connect2naga/go-examples/grpc/secure_grpc/proto"
	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/2/22 4:06 PM
Project : grpc-example
File : client.go
*/

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail the port 50052, error: %v ", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	res, err := c.SayHello(context.Background(), &chat.MessageReq{Body: "Hi this is Nagarjuna...."})
	if err != nil {
		log.Fatalf("error when calling sayHello, error:%v", err)
	}
	log.Printf("response from server : %s", res.Body)
}
