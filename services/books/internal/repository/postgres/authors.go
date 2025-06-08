package postgres

import (
	"context"
	"database/sql"

	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type (
	AuthorRepositoryImpl struct {
		dig.In
		DB *sql.DB
	}

	AuthorRepository interface {
		CreateAuthor(ctx context.Context, req domain.Author) (err error)
	}
)

func NewAuthorRepository(impl AuthorRepositoryImpl) AuthorRepository {
	return &impl
}

func (r *AuthorRepositoryImpl) CreateAuthor(ctx context.Context, req domain.Author) (err error) {
	arguments := []interface{}{
		req.AuthorID,
		req.Name,
		req.Biography,
	}

	_, err = r.DB.ExecContext(ctx, QueryCreateAuthor, arguments...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QueryCreateAuthor")
		return err
	}

	return nil
}
