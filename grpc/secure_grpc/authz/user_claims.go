// Package jwt_manager contains ...
package authz

import (
	"github.com/golang-jwt/jwt/v4"
)

/*
Author : Nagarjuna S
Date : 16-05-2022 00:18
Project : secure_grpc
File : user_claims.go
*/

type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}
