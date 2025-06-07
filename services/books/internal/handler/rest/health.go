package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	HealthControllerImpl struct {
		dig.In
	}

	HealthController interface {
		HealthCheck(ec echo.Context) (err error)
	}
)

func NewHealthHandler(impl HealthControllerImpl) HealthController {
	return &impl
}

func (c *HealthControllerImpl) HealthCheck(ec echo.Context) (err error) {
	return DefaultResult(ec, http.StatusOK)
}
