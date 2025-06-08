package postgres

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/books/internal/domain"
)

type (
	AuthorRepositoryImpl struct {
		dig.In
		DB *sql.DB
	}

	AuthorRepository interface {
		CreateAuthor(ctx context.Context, req domain.AuthorRequest) (err error)
		GetAuthorByID(ctx context.Context, authorID string) (author domain.Author, err error)
	}
)

func NewAuthorRepository(impl AuthorRepositoryImpl) AuthorRepository {
	return &impl
}

func (r *AuthorRepositoryImpl) CreateAuthor(ctx context.Context, req domain.AuthorRequest) (err error) {
	arguments := []interface{}{
		req.AuthorID,
		req.Name,
		req.Biography,
		req.CreatedBy,
	}

	_, err = r.DB.ExecContext(ctx, QueryCreateAuthor, arguments...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QueryCreateAuthor")
		return err
	}

	return nil
}

func (r *AuthorRepositoryImpl) GetAuthorByID(
	ctx context.Context, authorID string,
) (author domain.Author, err error) {
	rows, err := r.DB.QueryContext(ctx, QueryGetAuthorByID, authorID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while executing query QueryGetAuthorByID (authorID: %s)", authorID)
		return author, err
	}
	defer rows.Close()

	for rows.Next() {
		var scanner domain.AuthorScanner
		err = rows.Scan(
			&scanner.AuthorID,
			&scanner.Name,
			&scanner.Biography,
		)
		if err != nil {
			log.Ctx(ctx).Err(err).Msgf("while scanning row in QueryGetAuthorByID (authorID: %s)", authorID)
			return author, err
		}

		author = domain.Author{
			AuthorID:  scanner.AuthorID.String,
			Name:      scanner.Name.String,
			Biography: scanner.Biography.String,
		}
	}

	return author, nil
}
