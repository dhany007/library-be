package domain

import "database/sql"

type (
	Book struct {
		BookID      string   `json:"book_id"`
		Title       string   `json:"title"`
		ISBN        string   `json:"isbn"`
		Stock       int32    `json:"stock"`
		Description string   `json:"description"`
		Author      Author   `json:"author"`
		Category    Category `json:"category"`
	}

	BookScanner struct {
		BookID              sql.NullString `db:"book_id"`
		Title               sql.NullString `db:"title"`
		ISBN                sql.NullString `db:"isbn"`
		Stock               sql.NullInt32  `db:"stock"`
		Description         sql.NullString `db:"description"`
		AuthorID            sql.NullString `db:"author_id"`
		AuthorName          sql.NullString `db:"author_name"`
		Biography           sql.NullString `db:"author_bio"`
		CategoryID          sql.NullString `db:"category_id"`
		CategoryName        sql.NullString `db:"category_name"`
		CategoryDescription sql.NullString `db:"category_description"`
	}

	BookRequest struct {
		BookID      string
		Title       string `json:"title" valid:"required,minstringlength(3)"`
		ISBN        string `json:"isbn" valid:"required,minstringlength(10),maxstringlength(13)"`
		Stock       int32  `json:"stock" valid:"required,range(1|2147483647)"`
		Description string `json:"description"`
		AuthorID    string `json:"author_id" valid:"required,uuid"`
		CategoryID  string `json:"category_id" valid:"required,uuid"`
		CreatedBy   string
	}

	SearchBooksRequest struct {
		Title      string `json:"title"`
		ISBN       string `json:"isbn"`
		AuthorID   string `json:"author_id"`
		CategoryID string `json:"category_id"`
	}

	BorrowBookRequest struct {
		BookID string `json:"book_id" valid:"required,uuid"`
		UserID string `json:"user_id" valid:"required,uuid"`
	}
)

func (b *BookScanner) ToBook() Book {
	return Book{
		BookID:      b.BookID.String,
		Title:       b.Title.String,
		ISBN:        b.ISBN.String,
		Stock:       b.Stock.Int32,
		Description: b.Description.String,
		Author: Author{
			AuthorID:  b.AuthorID.String,
			Name:      b.AuthorName.String,
			Biography: b.Biography.String,
		},
		Category: Category{
			CategoryID:  b.CategoryID.String,
			Name:        b.CategoryName.String,
			Description: b.CategoryDescription.String,
		},
	}
}
