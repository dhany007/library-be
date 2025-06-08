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
	AuthorServiceImpl struct {
		dig.In
		AuthorRepository postgres.AuthorRepository
	}

	AuthorService interface {
		CreateAuthor(ctx context.Context, req domain.AuthorRequest) (author domain.Author, err error)
		GetAuthorByID(ctx context.Context, authorID string) (author domain.Author, err error)
	}
)

func NewAuthorService(impl AuthorServiceImpl) AuthorService {
	return &impl
}

func (s *AuthorServiceImpl) CreateAuthor(
	ctx context.Context, req domain.AuthorRequest,
) (author domain.Author, err error) {
	authorID := uuid.New().String()

	author = domain.Author{
		AuthorID:  authorID,
		Name:      req.Name,
		Biography: req.Biography,
	}

	err = s.AuthorRepository.CreateAuthor(ctx, author)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while AuthorRepository.CreateAuthor (authorID: %s)", authorID)
		return author, err
	}

	return author, nil
}

func (s *AuthorServiceImpl) GetAuthorByID(ctx context.Context, authorID string) (author domain.Author, err error) {
	author, err = s.AuthorRepository.GetAuthorByID(ctx, authorID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while AuthorRepository.GetAuthorByID (authorID: %s)", authorID)
		return author, err
	}

	if author.AuthorID == "" {
		err = domain.ErrAuthorNotFound
		log.Ctx(ctx).Err(err).Msgf("author not found (authorID: %s)", authorID)
		return author, err
	}

	return author, nil
}
