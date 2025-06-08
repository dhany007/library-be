package postgres

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/books/internal/domain"
)

type (
	CategoryRepositoryImpl struct {
		dig.In
		DB *sql.DB
	}

	CategoryRepository interface {
		CreateCategory(ctx context.Context, req domain.Category) (err error)
		GetCategoryByID(ctx context.Context, id string) (category domain.Category, err error)
	}
)

func NewCategoryRepository(impl CategoryRepositoryImpl) CategoryRepository {
	return &impl
}

func (r *CategoryRepositoryImpl) CreateCategory(ctx context.Context, req domain.Category) (err error) {
	arguments := []interface{}{
		req.CategoryID,
		req.Name,
		req.Description,
	}

	_, err = r.DB.ExecContext(ctx, QueryCreateCategory, arguments...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QueryCreateCategory")
		return err
	}

	return nil
}

func (r *CategoryRepositoryImpl) GetCategoryByID(
	ctx context.Context, id string,
) (category domain.Category, err error) {
	rows, err := r.DB.QueryContext(ctx, QueryGetCategoryByID, id)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while executing query QueryGetCategoryByID (id: %s)", id)
		return category, err
	}
	defer rows.Close()

	for rows.Next() {
		var scanner domain.CategoryScanner
		err = rows.Scan(
			&scanner.CategoryID,
			&scanner.Name,
			&scanner.Description,
		)
		if err != nil {
			log.Ctx(ctx).Err(err).Msgf("while scanning row for category (id: %s)", id)
			return category, err
		}

		category = domain.Category{
			CategoryID:  scanner.CategoryID.String,
			Name:        scanner.Name.String,
			Description: scanner.Description.String,
		}
	}

	return category, nil
}
