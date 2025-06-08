package domain

import "errors"

var (
	ErrInvalidAuthorID error = errors.New("invalid author ID")
	ErrAuthorNotFound  error = errors.New("author not found")
)
