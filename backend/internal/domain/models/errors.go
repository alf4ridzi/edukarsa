package models

import "errors"

var (
	ErrUsernameExist = errors.New("username already exist")
	ErrEmailExist    = errors.New("email already exist")
	ErrWrongPassword = errors.New("password is wrong")
	ErrForbidden     = errors.New("forbidden")
)
