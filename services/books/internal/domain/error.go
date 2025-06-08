package domain

import "errors"

var (
	ErrInvalidAuthorID         error = errors.New("invalid author ID")
	ErrAuthorNotFound          error = errors.New("author not found")
	ErrInvalidCategoryIDFormat error = errors.New("invalid category ID format")
	ErrCategoryNotFound        error = errors.New("category not found")
	ErrInvalidBookIDFormat     error = errors.New("invalid book ID format")
	ErrBookNotFound            error = errors.New("book not found")
	ErrISBNAlreadyExists       error = errors.New("ISBN already exists")
)
