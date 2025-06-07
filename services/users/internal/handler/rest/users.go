package rest

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/users/internal/domain"
	"github.com/dhany007/library-be/services/users/internal/services"
)

type (
	UsersControllerImpl struct {
		dig.In
		UserService services.UsersService
	}

	UsersController interface {
		Register(ec echo.Context) (err error)
	}
)

func NewUsersHandler(impl UsersControllerImpl) UsersController {
	return &impl
}

func (h *UsersControllerImpl) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req domain.RegisterRequest
	if err := c.Bind(&req); err != nil {
		log.Ctx(ctx).Err(err).Msg("while bind request body")
		return ResultError(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		log.Ctx(ctx).Err(err).Msg("while validate request body")
		return ResultError(c, http.StatusUnprocessableEntity, err)
	}

	user, err := h.UserService.Register(ctx, req)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while register user")
		return ResultError(c, http.StatusInternalServerError, err)
	}
	log.Ctx(ctx).Info().Msgf("user %s registered successfully", user.UserID)

	return ResultWithData(c, http.StatusCreated, user)
}
