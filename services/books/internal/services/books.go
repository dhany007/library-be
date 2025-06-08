package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/dhany007/library-be/services/books/internal/repository/postgres"
)

type (
	BookServiceImpl struct {
		dig.In
		CategoryRepository postgres.CategoryRepository
		AuthorRepository   postgres.AuthorRepository
		BookRepository     postgres.BookRepository
		TxRepository       postgres.TransactionRepository
	}

	BookService interface {
		CreateBook(ctx context.Context, req domain.BookRequest) (book domain.Book, err error)
		GetBookByID(ctx context.Context, id string) (book domain.Book, err error)
		SearchBooks(
			ctx context.Context, query domain.SearchBooksRequest,
		) (books []domain.Book, err error)
		BorrowBook(c context.Context, req domain.BorrowBookRequest) (err error)
		ReturnBook(ctx context.Context, req domain.BorrowBookRequest) (err error)
	}
)

func NewBookService(impl BookServiceImpl) BookService {
	return &impl
}

func (s *BookServiceImpl) CreateBook(
	ctx context.Context, req domain.BookRequest,
) (book domain.Book, err error) {
	bookID := uuid.New().String()

	category, err := s.CategoryRepository.GetCategoryByID(ctx, req.CategoryID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while CategoryRepository.GetCategoryByID (categoryID: %s)", req.CategoryID)
		return book, err
	}

	if category.CategoryID == "" {
		err = domain.ErrCategoryNotFound
		log.Ctx(ctx).Err(err).Msgf("while CategoryRepository.GetCategoryByID (categoryID: %s)", req.CategoryID)
		return book, err
	}

	author, err := s.AuthorRepository.GetAuthorByID(ctx, req.AuthorID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while AuthorRepository.GetAuthorByID (authorID: %s)", req.AuthorID)
		return book, err
	}

	if author.AuthorID == "" {
		err = domain.ErrAuthorNotFound
		log.Ctx(ctx).Err(err).Msgf("while AuthorRepository.GetAuthorByID (authorID: %s)", req.AuthorID)
		return book, err
	}

	req.BookID = bookID

	err = s.BookRepository.CreateBook(ctx, req)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while BookRepository.CreateBook (bookID: %s)", bookID)
		return book, err
	}

	book, err = s.BookRepository.GetBookByID(ctx, bookID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while BookRepository.GetBookByID (bookID: %s)", bookID)
		return book, err
	}

	return book, nil
}

func (s *BookServiceImpl) GetBookByID(
	ctx context.Context, id string,
) (book domain.Book, err error) {
	book, err = s.BookRepository.GetBookByID(ctx, id)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while BookRepository.GetBookByID (id: %s)", id)
		return book, err
	}

	if book.BookID == "" {
		err = domain.ErrBookNotFound
		log.Ctx(ctx).Err(err).Msgf("book not found (id: %s)", id)
		return book, err
	}

	return book, nil
}

func (s *BookServiceImpl) SearchBooks(
	ctx context.Context, query domain.SearchBooksRequest,
) (books []domain.Book, err error) {
	books, err = s.BookRepository.SearchBooks(ctx, query)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while BookRepository.SearchBooks (query: %+v)", query)
		return books, err
	}

	if len(books) == 0 {
		err = domain.ErrBookNotFound
		log.Ctx(ctx).Err(err).Msgf("no books found for query: %+v", query)
		return books, err
	}

	return books, nil
}

func (s *BookServiceImpl) BorrowBook(ctx context.Context, req domain.BorrowBookRequest) (err error) {
	tx, err := s.TxRepository.BeginTx()
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while TxRepository.BeginTx")
		return err
	}
	defer s.TxRepository.RollbackTx(tx)

	book, err := s.BookRepository.SelectBookForUpdate(ctx, tx, req.BookID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while BookRepository.SelectBookForUpdate (bookID: %s)", req.BookID)
		return err
	}

	if book.BookID == "" {
		err = domain.ErrBookNotFound
		log.Ctx(ctx).Err(err).Msgf("book not found (bookID: %s)", req.BookID)
		return err
	}

	if book.Stock <= 0 {
		err = domain.ErrStockBookEmpty
		log.Ctx(ctx).Err(err).Msgf("stock book is empty (bookID: %s)", req.BookID)
		return err
	}

	err = s.BookRepository.UpdateStockBook(ctx, tx, book.BookID, book.Stock-1)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while UpdateStockBook (bookID: %s)", req.BookID)
		return err
	}

	err = s.TxRepository.CommitTx(tx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while TxRepository.CommitTx")
		return err
	}

	return nil
}

func (s *BookServiceImpl) ReturnBook(ctx context.Context, req domain.BorrowBookRequest) (err error) {
	tx, err := s.TxRepository.BeginTx()
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while TxRepository.BeginTx")
		return err
	}
	defer s.TxRepository.RollbackTx(tx)

	book, err := s.BookRepository.SelectBookForUpdate(ctx, tx, req.BookID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while BookRepository.SelectBookForUpdate (bookID: %s)", req.BookID)
		return err
	}

	if book.BookID == "" {
		err = domain.ErrBookNotFound
		log.Ctx(ctx).Err(err).Msgf("book not found (bookID: %s)", req.BookID)
		return err
	}

	err = s.BookRepository.UpdateStockBook(ctx, tx, book.BookID, book.Stock+1)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while UpdateStockBook (bookID: %s)", req.BookID)
		return err
	}

	err = s.TxRepository.CommitTx(tx)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while TxRepository.CommitTx")
		return err
	}

	return nil
}
