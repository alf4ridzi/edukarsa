package models

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}
