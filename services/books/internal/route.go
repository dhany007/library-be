package internal

import (
	echo "github.com/labstack/echo/v4"

	"github.com/dhany007/library-be/services/books/internal/handler/rest"
)

const (
	ContextPath = "/v1/books/"

	HealthCheckPath = ContextPath + "health"

	CreateAuthor  = ContextPath + "authors"
	GetAuthorByID = ContextPath + "authors/:id"
)

func setRoute(
	e *echo.Echo,
	health rest.HealthController,
	author rest.AuthorController,
) {
	e.GET(HealthCheckPath, health.HealthCheck)

	e.POST(CreateAuthor, author.CreateAuthor)
	e.GET(GetAuthorByID, author.GetAuthorByID)
}
