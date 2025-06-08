package internal

import (
	echo "github.com/labstack/echo/v4"

	"github.com/dhany007/library-be/services/books/internal/handler/rest"
)

const (
	ContextPath = "/v1/"

	HealthCheckPath = ContextPath + "books/health"

	CreateAuthor  = ContextPath + "books/authors"
	GetAuthorByID = ContextPath + "books/authors/:id"

	CreateCategory  = ContextPath + "books/categories"
	GetCategoryByID = ContextPath + "books/categories/:id"

	CreateBook  = ContextPath + "books"
	GetBookByID = ContextPath + "books/:id"
)

func setRoute(
	e *echo.Echo,
	health rest.HealthController,
	author rest.AuthorController,
	category rest.CategoryController,
	book rest.BookController,
) {
	e.GET(HealthCheckPath, health.HealthCheck)

	e.POST(CreateAuthor, author.CreateAuthor)
	e.GET(GetAuthorByID, author.GetAuthorByID)

	e.POST(CreateCategory, category.CreateCategory)
	e.GET(GetCategoryByID, category.GetCategoryByID)

	e.POST(CreateBook, book.CreateBook)
	e.GET(GetBookByID, book.GetBookByID)
}
