package domain

import "errors"

var (
	ErrUserAlreadyExists error = errors.New("user already exists")
	ErrUserNotFound      error = errors.New("user not found")
	ErrPasswordNotMatch  error = errors.New("password not match")
)
