package postgres

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/books/internal/domain"
)

type (
	BookRepositoryImpl struct {
		dig.In
		DB *sql.DB
	}

	BookRepository interface {
		CreateBook(ctx context.Context, req domain.BookRequest) (err error)
		GetBookByID(ctx context.Context, bookID string) (book domain.Book, err error)
	}
)

func NewBookRepository(impl BookRepositoryImpl) BookRepository {
	return &impl
}

func (r *BookRepositoryImpl) CreateBook(ctx context.Context, req domain.BookRequest) (err error) {
	arguments := []interface{}{
		req.BookID,
		req.Title,
		req.ISBN,
		req.Stock,
		req.AuthorID,
		req.CategoryID,
		req.Description,
		req.CreatedBy,
	}

	_, err = r.DB.ExecContext(ctx, QueryCreateBook, arguments...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QueryCreateBook")
		return err
	}

	return nil
}

func (r *BookRepositoryImpl) GetBookByID(
	ctx context.Context, bookID string,
) (book domain.Book, err error) {
	rows, err := r.DB.QueryContext(ctx, QueryGetBookByID, bookID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while executing query QueryGetBookByID (bookID: %s)", bookID)
		return book, err
	}
	defer rows.Close()

	for rows.Next() {
		var scanner domain.BookScanner
		err = rows.Scan(
			&scanner.BookID,
			&scanner.Title,
			&scanner.ISBN,
			&scanner.Stock,
			&scanner.Description,
			&scanner.AuthorID,
			&scanner.AuthorName,
			&scanner.Biography,
			&scanner.CategoryID,
			&scanner.CategoryName,
			&scanner.CategoryDescription,
		)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("while scanning row in GetBookByID")
			return book, err
		}

		book = scanner.ToBook()
	}

	return book, nil
}
