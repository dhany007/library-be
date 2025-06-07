package infra

import (
	"time"

	echo "github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

type (
	AppCfg struct {
		Address      string        `envconfig:"ADDRESS" default:":9090" required:"true"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
		WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT" default:"10s"`
		Debug        bool          `envconfig:"DEBUG" default:"true"`
		Version      string        `envconfig:"VERSION" default:"v1"`
	}
)

func NewEcho(cfg *AppCfg) *echo.Echo {
	e := echo.New()

	e.Use(em.Recover())
	e.Use(em.CORS())
	e.Use(em.Gzip())
	e.Use(em.RequestID())

	e.HideBanner = true
	return e
}
