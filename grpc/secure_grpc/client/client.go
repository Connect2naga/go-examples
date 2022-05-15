// Package client contains ...
package client

import (
	"context"
	"time"

	chat "github.com/connect2naga/go-examples/grpc/secure_grpc/proto"
	"google.golang.org/grpc"
)

/*
Author : Nagarjuna S
Date : 2/2/22 4:06 PM
Project : grpc-example
File : client.go
*/

type AuthClient struct {
	service  chat.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := chat.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &chat.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}
