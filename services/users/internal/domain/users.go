package domain

import (
	"database/sql"
)

type (
	User struct {
		UserID       string `json:"user_id"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		PasswordHash string `json:"-"`
		Role         string `json:"role"`
	}

	UserScanner struct {
		UserID       sql.NullString `db:"id"`
		Email        sql.NullString `db:"email"`
		Name         sql.NullString `db:"name"`
		Role         sql.NullString `db:"role"`
		PasswordHash sql.NullString `db:"password_hash"`
	}

	RegisterRequest struct {
		UserID   string `json:"user_id" `
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required,minstringlength(8)"`
		Name     string `json:"name" valid:"required,minstringlength(3)"`
		Role     string `json:"role" valid:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required,minstringlength(8)"`
	}

	Authorization struct {
		UserID string `json:"user_id"`
		Role   string `json:"role"`
		Email  string `json:"email"`
	}
)
