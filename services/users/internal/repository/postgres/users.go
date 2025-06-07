package postgres

import (
	"context"
	"database/sql"

	"go.uber.org/dig"

	"github.com/dhany007/library-be/services/users/internal/domain"
	"github.com/rs/zerolog/log"
)

type (
	UsersRepositoryImpl struct {
		dig.In
		DB *sql.DB
	}

	UsersRepository interface {
		CreateUser(ctx context.Context, req domain.RegisterRequest) (err error)
		GetUserByEmail(ctx context.Context, email string) (user domain.User, err error)
	}
)

func NewUsersRepository(impl UsersRepositoryImpl) UsersRepository {
	return &impl
}

func (r *UsersRepositoryImpl) CreateUser(ctx context.Context, req domain.RegisterRequest) (err error) {
	arguments := []interface{}{
		req.UserID,
		req.Email,
		req.Name,
		req.Password,
		req.Role,
	}

	_, err = r.DB.ExecContext(ctx, QueryCreateUser, arguments...)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QueryCreateUser")
		return err
	}

	return nil
}

func (r *UsersRepositoryImpl) GetUserByEmail(
	ctx context.Context, email string,
) (user domain.User, err error) {
	rows, err := r.DB.QueryContext(ctx, QueryGetUserByEmail, email)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("while executing query QueryGetUserByEmail")
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		var scanner domain.UserScanner

		err = rows.Scan(
			&scanner.UserID,
			&scanner.Email,
			&scanner.Name,
			&scanner.PasswordHash,
			&scanner.Role,
		)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("while scanning row in QueryGetUserByEmail")
			return user, err
		}

		user = domain.User{
			UserID:       scanner.UserID.String,
			Email:        scanner.Email.String,
			PasswordHash: scanner.PasswordHash.String,
			Name:         scanner.Name.String,
			Role:         scanner.Role.String,
		}
	}

	return user, nil
}
