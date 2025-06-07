package services

import (
	"context"

	"github.com/dhany007/library-be/services/users/internal/domain"
	"github.com/dhany007/library-be/services/users/internal/repository/postgres"
	"github.com/dhany007/library-be/services/users/pkg/bcrypt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type (
	UsersServiceImpl struct {
		dig.In
		UsersRepository postgres.UsersRepository
	}

	UsersService interface {
		Register(ctx context.Context, req domain.RegisterRequest) (user domain.User, err error)
	}
)

func NewUsersService(impl UsersServiceImpl) UsersService {
	return &impl
}

func (s *UsersServiceImpl) Register(
	ctx context.Context, req domain.RegisterRequest,
) (user domain.User, err error) {
	userID := uuid.New().String()

	userExists, err := s.UsersRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while UsersRepository.GetUserByEmail (email: %s)", req.Email)
		return user, err
	}

	if userExists.UserID != "" {
		log.Ctx(ctx).Err(domain.ErrUserAlreadyExists).Msgf("user with email %s already exists", req.Email)
		return user, domain.ErrUserAlreadyExists
	}

	passwordHash := bcrypt.HashPass(req.Password)
	req.Password = passwordHash
	req.UserID = userID

	err = s.UsersRepository.CreateUser(ctx, req)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while UsersRepository.CreateUser (email: %s)", req.Email)
		return user, err
	}

	user = domain.User{
		UserID: userID,
		Email:  req.Email,
		Name:   req.Name,
		Role:   req.Role,
	}

	return user, nil
}
