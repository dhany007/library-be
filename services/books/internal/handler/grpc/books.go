package grpc

import (
	"context"

	bookpb "github.com/dhany007/library-be/proto/bookpb"
	"github.com/dhany007/library-be/services/books/internal/services"
)

type BookGRPCHandler struct {
	bookpb.UnimplementedBookServiceServer
	svc services.BookService
}

func NewBookGRPCHandler(svc services.BookService) *BookGRPCHandler {
	return &BookGRPCHandler{svc: svc}
}

func (h *BookGRPCHandler) CreateBook(ctx context.Context, req *bookpb.CreateBookRequest) (*bookpb.BookResponse, error) {
	return &bookpb.BookResponse{}, nil
}
