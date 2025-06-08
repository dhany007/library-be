package rest

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/dhany007/library-be/services/books/internal/services"
)

type (
	BookControllerImpl struct {
		dig.In
		BookService services.BookService
	}

	BookController interface {
		CreateBook(c echo.Context) (err error)
		GetBookByID(c echo.Context) (err error)
		SearchBooks(c echo.Context) (err error)
	}
)

func NewBookHandler(impl BookControllerImpl) BookController {
	return &impl
}

func (h *BookControllerImpl) CreateBook(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req domain.BookRequest
	if err := c.Bind(&req); err != nil {
		log.Ctx(ctx).Err(err).Msgf("while bind request body (title: %s)", req.Title)
		return ResultError(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(&req); err != nil {
		log.Ctx(ctx).Err(err).Msgf("while validate request body (title: %s)", req.Title)
		return ResultError(c, http.StatusUnprocessableEntity, err)
	}

	book, err := h.BookService.CreateBook(ctx, req)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while create book (title: %s)", req.Title)
		return ResultError(c, http.StatusInternalServerError, err)
	}

	log.Ctx(ctx).Info().Msgf("book %s created successfully", book.BookID)

	return ResultWithData(c, http.StatusCreated, book)
}

func (h *BookControllerImpl) GetBookByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	bookID := c.Param("id")
	if isUUID := govalidator.IsUUID(bookID); !isUUID {
		err = domain.ErrInvalidBookIDFormat
		log.Ctx(ctx).Err(err).Msgf("invalid book ID format (id: %s)", bookID)
		return ResultError(c, http.StatusBadRequest, err)
	}

	book, err := h.BookService.GetBookByID(ctx, bookID)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while get book by id (id: %s)", bookID)
		return ResultError(c, http.StatusInternalServerError, err)
	}

	log.Ctx(ctx).Info().Msgf("book %s retrieved successfully", book.BookID)

	return ResultWithData(c, http.StatusOK, book)
}

func (h *BookControllerImpl) SearchBooks(c echo.Context) (err error) {
	ctx := c.Request().Context()

	query := domain.SearchBooksRequest{
		Title:      c.QueryParam("title"),
		ISBN:       c.QueryParam("isbn"),
		AuthorID:   c.QueryParam("author_id"),
		CategoryID: c.QueryParam("category_id"),
	}

	books, err := h.BookService.SearchBooks(ctx, query)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while searching books with query: %s", query)
		return ResultError(c, http.StatusInternalServerError, err)
	}

	log.Ctx(ctx).Info().Msgf("found %d books for query: %s", len(books), query)

	return ResultWithData(c, http.StatusOK, books)
}
