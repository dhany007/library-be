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
	}

	BookService interface {
		CreateBook(ctx context.Context, req domain.BookRequest) (book domain.Book, err error)
		GetBookByID(ctx context.Context, id string) (book domain.Book, err error)
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
