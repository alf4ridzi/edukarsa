package utils

import (
	"edukarsa-backend/internal/config"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAccessClaims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type JwtRefreshclaims struct {
	jwt.RegisteredClaims
}

func CreateAccessToken(id uint, role string) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(config.AppConfig.AccessTokenExpired))
	claims := &JwtAccessClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(id), 10),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.AccessSecret))
}

func CreateRefreshToken(id uint) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(config.AppConfig.RefreshTokenExpired))
	claims := &JwtRefreshclaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(id), 10),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.RefreshSecret))
}

func ValidateAccessToken(tokenJwt string) (*JwtAccessClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenJwt,
		&JwtAccessClaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(config.AppConfig.AccessSecret), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*JwtAccessClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

func ValidateRefreshToken(tokenJwt string) (*JwtRefreshclaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenJwt,
		&JwtRefreshclaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(config.AppConfig.RefreshSecret), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*JwtRefreshclaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
