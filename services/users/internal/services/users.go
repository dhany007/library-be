package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/users/internal/domain"
	"github.com/dhany007/library-be/services/users/internal/repository/postgres"
	"github.com/dhany007/library-be/services/users/pkg/bcrypt"
	"github.com/dhany007/library-be/services/users/pkg/jwt"
)

type (
	UsersServiceImpl struct {
		dig.In
		UsersRepository postgres.UsersRepository
	}

	UsersService interface {
		Register(ctx context.Context, req domain.RegisterRequest) (user domain.User, err error)
		Login(ctx context.Context, req domain.LoginRequest) (token string, err error)
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

func (s *UsersServiceImpl) Login(
	ctx context.Context, req domain.LoginRequest,
) (token string, err error) {

	user, err := s.UsersRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while UsersRepository.GetUserByEmail (email: %s)", req.Email)
		return token, err
	}

	if user.UserID == "" {
		log.Ctx(ctx).Err(domain.ErrUserNotFound).Msgf("user with email %s not found", req.Email)
		return token, domain.ErrUserNotFound
	}

	if !bcrypt.ComparePassword([]byte(user.PasswordHash), []byte(req.Password)) {
		log.Ctx(ctx).Err(domain.ErrPasswordNotMatch).Msgf("invalid password for user with email %s", req.Email)
		return token, domain.ErrPasswordNotMatch
	}

	duration := time.Duration(domain.DurationTTLTokenMinutes * time.Minute)
	token, err = jwt.GenerateToken(duration, domain.Authorization{
		UserID: user.UserID,
		Role:   user.Role,
		Email:  user.Email,
	})
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("while generating token for user with email %s", req.Email)
		return token, err
	}

	return token, nil
}
