package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
		SearchBooks(
			ctx context.Context, query domain.SearchBooksRequest,
		) (books []domain.Book, err error)
		SelectBookForUpdate(
			ctx context.Context, tx *sql.Tx, bookID string,
		) (book domain.Book, err error)
		UpdateStockBook(ctx context.Context, tx *sql.Tx, bookID string, stok int32) error
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

func (r *BookRepositoryImpl) SearchBooks(
	ctx context.Context, query domain.SearchBooksRequest,
) (books []domain.Book, err error) {
	var (
		filters []string
		args    []interface{}
		idx     = 1
	)

	if query.Title != "" {
		filters = append(filters, fmt.Sprintf("LOWER(b.title) ILIKE $%d", idx))
		args = append(args, "%"+strings.ToLower(query.Title)+"%")
		idx++
	}

	if query.ISBN != "" {
		filters = append(filters, fmt.Sprintf("LOWER(b.isbn) = $%d", idx))
		args = append(args, strings.ToLower(query.ISBN))
		idx++
	}

	if query.AuthorID != "" {
		filters = append(filters, fmt.Sprintf("LOWER(a.id) = $%d", idx))
		args = append(args, strings.ToLower(query.AuthorID))
		idx++
	}

	if query.CategoryID != "" {
		filters = append(filters, fmt.Sprintf("LOWER(c.id) = $%d", idx))
		args = append(args, strings.ToLower(query.CategoryID))
		idx++
	}

	q := QuerySearchBooks
	if len(filters) > 0 {
		q += " AND " + strings.Join(filters, " AND ")
	}

	rows, err := r.DB.QueryContext(ctx, q, args...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QuerySearchBooks")
		return books, err
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
			log.Ctx(ctx).Err(err).Msg("while scanning row in SearchBooks")
			return books, err
		}

		book := scanner.ToBook()
		if book.BookID != "" {
			books = append(books, book)
		}
	}

	return books, nil
}

func (r *BookRepositoryImpl) SelectBookForUpdate(
	ctx context.Context, tx *sql.Tx, bookID string,
) (book domain.Book, err error) {
	rows, err := tx.QueryContext(ctx, QuerySelectBookForUpdate, bookID)
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
		)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("while scanning row in GetBookByID")
			return book, err
		}

		book = scanner.ToBook()
	}

	return book, nil
}

func (r BookRepositoryImpl) UpdateStockBook(ctx context.Context, tx *sql.Tx, bookID string, stok int32) error {
	_, err := tx.ExecContext(ctx, QueryUpdateStockBook, bookID, stok)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while ExecContext QueryUpdateStockBook (bookID: %s)", bookID)
		return err
	}

	return nil
}
