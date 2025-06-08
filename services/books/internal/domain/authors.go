package domain

import "database/sql"

type (
	Author struct {
		AuthorID  string `json:"author_id"`
		Name      string `json:"name"`
		Biography string `json:"biography"`
	}

	AuthorScanner struct {
		AuthorID  sql.NullString `db:"id"`
		Name      sql.NullString `db:"name"`
		Biography sql.NullString `db:"bio"`
	}

	AuthorRequest struct {
		AuthorID  string
		Name      string `json:"name" valid:"required,minstringlength(3)"`
		Biography string `json:"biography"`
		CreatedBy string
	}
)
