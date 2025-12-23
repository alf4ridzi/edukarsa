package utils

import "testing"

func TestRandomLetter(t *testing.T) {
	letter := GenerateRandomString(7)
	t.Log(letter)
}
