// Package cmd contains ...
package main

import (
	"github.com/connect2naga/go-examples/grpc/secure_grpc/authz"
	chat "github.com/connect2naga/go-examples/grpc/secure_grpc/proto"
	"github.com/connect2naga/go-examples/grpc/secure_grpc/server"
	"github.com/connect2naga/go-examples/grpc/secure_grpc/store"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

/*
Author : Nagarjuna S
Date : 16-05-2022 01:53
Project : secure_grpc
File : main.go
*/

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to open port :50052 for listen, error: %v", err)
	}

	userStore := store.NewInMemoryUserStore()

	err = seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	jwtManager := authz.NewJWTManager("1234", 180*time.Second)
	interceptor := authz.NewAuthInterceptor(jwtManager, authz.AccessibleRoles())
	authServer := server.NewAuthServer(userStore, jwtManager)

	s := server.Server{}
	gServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))

	chat.RegisterChatServiceServer(gServer, &s)
	chat.RegisterAuthServiceServer(gServer, authServer)

	if err := gServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serv GRPC Server over port 50052, %v", err)
	}
}

func createUser(userStore store.UserStore, username, password, role string) error {
	user, err := authz.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore store.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}
