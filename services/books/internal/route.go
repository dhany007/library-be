package internal

import (
	echo "github.com/labstack/echo/v4"

	"github.com/dhany007/library-be/services/books/internal/handler/rest"
)

const (
	ContextPath = "/v1/books/"

	HealthCheckPath = ContextPath + "health"
)

func setRoute(e *echo.Echo, h rest.HealthController) {
	e.GET(HealthCheckPath, h.HealthCheck)
}
