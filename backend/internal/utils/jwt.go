package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func CreateAccessToken(id uint, role string, secret string, expired int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expired))
	claims := &JwtClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(id), 10),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateRefreshToken(id uint, secret string, expired int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expired))
	claims := &JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(id), 10),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// func IsAuthorized(tokenJwt string, secret string) (bool, error) {

// }
