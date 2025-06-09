package grpc

import (
	"context"
	"errors"

	bookpb "github.com/dhany007/library-be/proto/bookpb"
	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/dhany007/library-be/services/books/internal/services"
	"github.com/rs/zerolog/log"
)

type BookGRPCHandler struct {
	bookpb.UnimplementedBookServiceServer
	svc services.BookService
}

func NewBookGRPCHandler(svc services.BookService) *BookGRPCHandler {
	return &BookGRPCHandler{svc: svc}
}

func (h *BookGRPCHandler) CreateBook(ctx context.Context, req *bookpb.CreateBookRequest) (out *bookpb.BookResponse, err error) {
	if req == nil {
		err = errors.New("request is required")
		log.Ctx(ctx).Err(err).Msgf("while validate request (req: %v)", req)
		return out, err
	}

	book, err := h.svc.CreateBook(ctx, domain.BookRequest{
		Title:       req.Title,
		ISBN:        req.Isbn,
		Stock:       req.Stock,
		Description: req.Description,
		AuthorID:    req.AuthorId,
		CategoryID:  req.CategoryId,
		CreatedBy:   req.CreatedBy,
	})

	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while svc.CreateBook (bookID: %v)", req)
		return out, err
	}

	return &bookpb.BookResponse{
		BookId:      book.BookID,
		Title:       book.Title,
		Isbn:        book.ISBN,
		Description: book.Description,
		Author: &bookpb.AuthorResponse{
			AuthorId:   book.Author.AuthorID,
			AuthorName: book.Author.Name,
			Biography:  book.Author.Biography,
		},
		Category: &bookpb.CategoryResponse{
			CategoryId:          book.Category.CategoryID,
			CategoryName:        book.Category.Name,
			CategoryDescription: book.Category.Description,
		},
	}, nil
}

func (h *BookGRPCHandler) BorrowBook(ctx context.Context, req *bookpb.BorrowBookRequest) (out *bookpb.BorrowBookResponse, err error) {
	if req == nil {
		err = errors.New("request is required")
		log.Ctx(ctx).Err(err).Msgf("while validate request (req: %v)", req)
		return out, err
	}

	err = h.svc.BorrowBook(ctx, domain.BorrowBookRequest{
		UserID: req.UserId,
		BookID: req.BookId,
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while svc.BorrowBook (bookID: %v)", req)
		return &bookpb.BorrowBookResponse{
			Success: false,
			Err:     err.Error(),
		}, nil
	}

	return &bookpb.BorrowBookResponse{
		Success: true,
		Err:     "",
	}, nil
}

func (h *BookGRPCHandler) ReturnBook(ctx context.Context, req *bookpb.BorrowBookRequest) (out *bookpb.BorrowBookResponse, err error) {
	if req == nil {
		err = errors.New("request is required")
		log.Ctx(ctx).Err(err).Msgf("while validate request (req: %v)", req)
		return out, err
	}

	err = h.svc.ReturnBook(ctx, domain.BorrowBookRequest{
		UserID: req.UserId,
		BookID: req.BookId,
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while svc.ReturnBook (bookID: %v)", req.BookId)
		return &bookpb.BorrowBookResponse{
			Success: false,
			Err:     err.Error(),
		}, nil
	}

	return &bookpb.BorrowBookResponse{
		Success: true,
		Err:     "",
	}, nil
}

func (h *BookGRPCHandler) GetBookByID(ctx context.Context, req *bookpb.GetBookByIDRequest) (out *bookpb.BookResponse, err error) {
	if req == nil || req.BookId == "" {
		err = domain.ErrInvalidBookIDFormat
		log.Ctx(ctx).Err(err).Msgf("while validaate bookID (req: %v)", req)
		return out, nil
	}

	book, err := h.svc.GetBookByID(ctx, req.BookId)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while svc.GetBookByID (bookID: %v)", req)
		return out, err
	}

	return &bookpb.BookResponse{
		BookId:      book.BookID,
		Title:       book.Title,
		Isbn:        book.ISBN,
		Description: book.Description,
		Author: &bookpb.AuthorResponse{
			AuthorId:   book.Author.AuthorID,
			AuthorName: book.Author.Name,
			Biography:  book.Author.Biography,
		},
		Category: &bookpb.CategoryResponse{
			CategoryId:          book.Category.CategoryID,
			CategoryName:        book.Category.Name,
			CategoryDescription: book.Category.Description,
		},
	}, nil
}
