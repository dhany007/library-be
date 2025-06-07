package internal

import (
	echo "github.com/labstack/echo/v4"

	"github.com/dhany007/library-be/services/users/internal/handler/rest"
)

const (
	ContextPath = "/v1/users/"

	HealthCheckPath = ContextPath + "health"

	RegisterPath = ContextPath + "register"
	LoginPath    = ContextPath + "login"
)

func setRoute(e *echo.Echo, health rest.HealthController, user rest.UsersController) {
	e.GET(HealthCheckPath, health.HealthCheck)

	e.POST(RegisterPath, user.Register)
	e.POST(LoginPath, user.Login)
}
