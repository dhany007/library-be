package rest

import (
	"net/http"

	"go.uber.org/dig"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/dhany007/library-be/services/books/internal/services"
)

type (
	AuthorControllerImpl struct {
		dig.In
		AuthorService services.AuthorService
	}

	AuthorController interface {
		CreateAuthor(c echo.Context) (err error)
	}
)

func NewAuthorHandler(impl AuthorControllerImpl) AuthorController {
	return &impl
}

func (h *AuthorControllerImpl) CreateAuthor(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req domain.AuthorRequest
	if err := c.Bind(&req); err != nil {
		log.Ctx(ctx).Err(err).Msgf("while bind request body (name: %s)", req.Name)
		return ResultError(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		log.Ctx(ctx).Err(err).Msgf("while validate request body (name: %s)", req.Name)
		return ResultError(c, http.StatusUnprocessableEntity, err)
	}

	author, err := h.AuthorService.CreateAuthor(ctx, req)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while create author (name: %s)", req.Name)
		return ResultError(c, http.StatusInternalServerError, err)
	}

	log.Ctx(ctx).Info().Msgf("author %s created successfully", author.AuthorID)

	return ResultWithData(c, http.StatusCreated, author)
}
