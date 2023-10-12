package dto

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	jwt.RegisteredClaims
	Name  string
	Admin bool
	Exp   int64
}
