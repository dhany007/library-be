package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/dhany007/library-be/services/users/internal/domain"
	"github.com/dhany007/library-be/services/users/pkg/env"
)

func GenerateToken(ttl time.Duration, auth domain.Authorization) (token string, err error) {
	now := time.Now().Unix()
	exp := time.Now().Add(ttl * time.Minute).Unix()

	key := env.GetEnv(domain.EnvKeyJwt, "secretkey")

	claims := jwt.MapClaims{
		"id":    auth.UserID,
		"email": auth.Email,
		"role":  auth.Role,
		"iat":   now,
		"exp":   exp,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = parseToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return token, nil
}
