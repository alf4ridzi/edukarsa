package utils

import (
	"edukarsa-backend/internal/domain/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *models.User, secret string, expired int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expired))
	claims := &models.JwtClaims{
		Role: user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
