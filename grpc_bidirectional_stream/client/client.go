// Package client contains ...
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	bi_stream "github.com/connect2naga/go-examples/grpc_bidirectional_stream/proto"
	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/4/22 2:00 PM
Project : go-examples
File : client.go
*/

func main() {
	rand.Seed(time.Now().Unix())

	// dail server
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := bi_stream.NewChatBiSteamServiceClient(conn)
	stream, err := client.SayHello(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	// send 10 messages and closes
	go func() {
		for i := 1; i <= 10; i++ {
			// generates random number and sends it to stream
			rnd := int32(rand.Intn(i))
			req := bi_stream.MessageReq{Body: fmt.Sprintf("Sent Message : %d", rnd)}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("Send to Server ===> %s", req.Body)
			time.Sleep(time.Millisecond * 1000)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("got from Server <===== %s", resp.Body)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("finished with")
}
