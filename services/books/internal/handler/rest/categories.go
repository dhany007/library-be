package rest

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/dhany007/library-be/services/books/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type (
	CategoryControllerImpl struct {
		dig.In
		CategoryService services.CategoryService
	}

	CategoryController interface {
		CreateCategory(c echo.Context) (err error)
		GetCategoryByID(c echo.Context) (err error)
	}
)

func NewCategoryHandler(impl CategoryControllerImpl) CategoryController {
	return &impl
}

func (h *CategoryControllerImpl) CreateCategory(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req domain.CategoryRequest
	if err := c.Bind(&req); err != nil {
		log.Ctx(ctx).Err(err).Msgf("while bind request body (name: %s)", req.Name)
		return ResultError(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		log.Ctx(ctx).Err(err).Msgf("while validate request body (name: %s)", req.Name)
		return ResultError(c, http.StatusUnprocessableEntity, err)
	}

	category, err := h.CategoryService.CreateCategory(ctx, req)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while create category (name: %s)", req.Name)
		return ResultError(c, http.StatusInternalServerError, err)
	}

	log.Ctx(ctx).Info().Msgf("category %s created successfully", category.CategoryID)

	return ResultWithData(c, http.StatusCreated, category)
}

func (h *CategoryControllerImpl) GetCategoryByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	categoryID := c.Param("id")
	if isUUID := govalidator.IsUUID(categoryID); !isUUID {
		err = domain.ErrInvalidCategoryIDFormat
		log.Ctx(ctx).Err(err).Msgf("invalid category ID format (id: %s)", categoryID)
		return ResultError(c, http.StatusBadRequest, err)
	}

	category, err := h.CategoryService.GetCategoryByID(ctx, categoryID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while get category by ID (id: %s)", categoryID)
		if err == domain.ErrCategoryNotFound {
			return ResultError(c, http.StatusNotFound, err)
		}

		return ResultError(c, http.StatusInternalServerError, err)
	}

	log.Ctx(ctx).Info().Msgf("category %s retrieved successfully", category.CategoryID)

	return ResultWithData(c, http.StatusOK, category)
}
