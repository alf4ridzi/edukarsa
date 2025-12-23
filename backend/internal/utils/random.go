package utils

import (
	"math/rand/v2"
)

const alphaNumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphaNumeric[rand.IntN(len(alphaNumeric))]
	}

	return string(b)
}
