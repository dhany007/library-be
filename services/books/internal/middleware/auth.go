package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/dhany007/library-be/services/books/internal/domain"
	"github.com/dhany007/library-be/services/users/pkg/env"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	key := env.GetEnv(domain.EnvKeyJwt, "")
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid Authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		now := time.Now().Unix()
		expired := claims["exp"].(float64)

		if now >= int64(expired) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		c.Set("id", claims["id"])
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])

		return next(c)
	}
}

func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, ok := c.Get("role").(string)
		if !ok || role != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "Admin access required")
		}
		return next(c)
	}
}
