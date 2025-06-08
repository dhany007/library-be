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

	CreateCategory  = ContextPath + "categories"
	GetCategoryByID = ContextPath + "categories/:id"
)

func setRoute(
	e *echo.Echo,
	health rest.HealthController,
	author rest.AuthorController,
	category rest.CategoryController,
) {
	e.GET(HealthCheckPath, health.HealthCheck)

	e.POST(CreateAuthor, author.CreateAuthor)
	e.GET(GetAuthorByID, author.GetAuthorByID)

	e.POST(CreateCategory, category.CreateCategory)
	e.GET(GetCategoryByID, category.GetCategoryByID)
}
