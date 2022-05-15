// Package client_cmd contains ...
package main

import (
	"context"
	"github.com/connect2naga/go-examples/grpc/secure_grpc/client"
	chat "github.com/connect2naga/go-examples/grpc/secure_grpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

/*
Author : Nagarjuna S
Date : 16-05-2022 01:58
Project : secure_grpc
File : main.go
*/

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail the port 50052, error: %v ", err)
	}
	defer conn.Close()

	authClient := client.NewAuthClient(conn, "admin1", "secret")
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), 180*time.Second)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	conn2, err := grpc.Dial(
		":50052",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	c := chat.NewChatServiceClient(conn2)
	res, err := c.SayHello(context.Background(), &chat.MessageReq{Body: "Hi this is Nagarjuna...."})
	if err != nil {
		log.Fatalf("error when calling sayHello, error:%v", err)
	}
	log.Printf("response from server : %s", res.Body)
}

func authMethods() map[string]bool {
	const experimentPath = "/exp/"

	return map[string]bool{
		experimentPath + "start":  true,
		experimentPath + "stop":   true,
		experimentPath + "status": true,
	}
}
