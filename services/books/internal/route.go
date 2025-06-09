package internal

import (
	echo "github.com/labstack/echo/v4"

	"github.com/dhany007/library-be/services/books/internal/handler/rest"
	em "github.com/dhany007/library-be/services/books/internal/middleware"
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
	SearchBooks = ContextPath + "books/search"

	BorrowBook = ContextPath + "books/borrow"
	ReturnBook = ContextPath + "books/return"
)

func setRoute(
	e *echo.Echo,
	health rest.HealthController,
	author rest.AuthorController,
	category rest.CategoryController,
	book rest.BookController,
) {
	e.GET(HealthCheckPath, health.HealthCheck)

	e.POST(CreateAuthor, em.AuthMiddleware(em.AdminOnlyMiddleware(author.CreateAuthor)))
	e.GET(GetAuthorByID, em.AuthMiddleware(author.GetAuthorByID))

	e.POST(CreateCategory, em.AuthMiddleware(em.AdminOnlyMiddleware(category.CreateCategory)))
	e.GET(GetCategoryByID, em.AuthMiddleware(category.GetCategoryByID))

	e.POST(CreateBook, em.AuthMiddleware(em.AdminOnlyMiddleware(book.CreateBook)))
	e.GET(GetBookByID, em.AuthMiddleware(book.GetBookByID))
	e.GET(SearchBooks, em.AuthMiddleware(book.SearchBooks))

	e.POST(BorrowBook, em.AuthMiddleware(book.BorrowBook))
	e.POST(ReturnBook, em.AuthMiddleware(book.ReturnBook))
}
