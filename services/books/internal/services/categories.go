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
	CategoryServiceImpl struct {
		dig.In
		CategoryRepository postgres.CategoryRepository
	}

	CategoryService interface {
		CreateCategory(ctx context.Context, req domain.CategoryRequest) (category domain.Category, err error)
		GetCategoryByID(ctx context.Context, id string) (category domain.Category, err error)
	}
)

func NewCategoryService(impl CategoryServiceImpl) CategoryService {
	return &impl
}

func (s *CategoryServiceImpl) CreateCategory(
	ctx context.Context, req domain.CategoryRequest,
) (category domain.Category, err error) {
	categoryID := uuid.New().String()

	category = domain.Category{
		CategoryID:  categoryID,
		Name:        req.Name,
		Description: req.Description,
	}

	err = s.CategoryRepository.CreateCategory(ctx, category)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while CategoryRepository.CreateCategory (categoryID: %s)", categoryID)
		return category, err
	}

	return category, nil
}

func (s *CategoryServiceImpl) GetCategoryByID(
	ctx context.Context, id string,
) (category domain.Category, err error) {
	category, err = s.CategoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while CategoryRepository.GetCategoryByID (id: %s)", id)
		return category, err
	}

	if category.CategoryID == "" {
		err = domain.ErrCategoryNotFound
		log.Ctx(ctx).Err(err).Msgf("category not found (id: %s)", id)
		return category, err
	}

	return category, nil
}
