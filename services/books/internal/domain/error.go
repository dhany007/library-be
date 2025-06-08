package domain

import "errors"

var (
	ErrInvalidAuthorID         error = errors.New("invalid author ID")
	ErrAuthorNotFound          error = errors.New("author not found")
	ErrInvalidCategoryIDFormat error = errors.New("invalid category ID format")
	ErrCategoryNotFound        error = errors.New("category not found")
)
