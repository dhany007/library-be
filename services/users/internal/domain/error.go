package domain

import "errors"

var (
	ErrUserAlreadyExists error = errors.New("user already exists")
)
