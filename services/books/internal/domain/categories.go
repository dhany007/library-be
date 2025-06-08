package domain

import "database/sql"

type (
	Category struct {
		CategoryID  string `json:"category_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	CategoryScanner struct {
		CategoryID  sql.NullString `db:"id"`
		Name        sql.NullString `db:"name"`
		Description sql.NullString `db:"description"`
	}

	CategoryRequest struct {
		CategoryID  string
		Name        string `json:"name" valid:"required,minstringlength(3)"`
		Description string `json:"description"`
		CreatedBy   string
	}
)
