// Package server_streaming contains ...
package main

import (
	"context"
	"io"
	"log"

	msg_stream "github.com/connect2naga/go-examples/grpc_server_stream/proto"

	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/4/22 10:06 AM
Project : grpc-example
File : client.go
*/

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dail the port 50052, error: %v ", err)
	}
	defer conn.Close()

	c := msg_stream.NewChatSteamServiceClient(conn)

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

	<-done
	log.Printf("all the messages recieved from server....")
}
