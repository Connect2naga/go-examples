// Package main contains ...
package server

import (
	"context"
	jwt_manager "github.com/connect2naga/go-examples/grpc/secure_grpc/authz"

	chat "github.com/connect2naga/go-examples/grpc/secure_grpc/proto"
	"github.com/connect2naga/go-examples/grpc/secure_grpc/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
Author : Nagarjuna S
Date : 16-05-2022 01:16
Project : secure_grpc
File : auth_server.go
*/

type AuthServer struct {
	userStore  store.UserStore
	jwtManager *jwt_manager.JWTManager
	chat.UnimplementedAuthServiceServer
}

func NewAuthServer(userStore store.UserStore, jwtManager *jwt_manager.JWTManager) *AuthServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}

func (server *AuthServer) Login(ctx context.Context, req *chat.LoginRequest) (*chat.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &chat.LoginResponse{AccessToken: token}
	return res, nil
}
