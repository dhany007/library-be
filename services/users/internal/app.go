package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/dhany007/library-be/services/users/internal/infra"
	"github.com/dhany007/library-be/services/users/pkg/di"
)

var exitSigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

func Start() {
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, exitSigs...)

	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		if err := di.Invoke(startApp); err != nil {
			log.Fatal().Msgf("startApp: %s", err.Error())
		}
	}()
	<-exitCh

	if err := di.Invoke(gracefulShutdown); err != nil {
		log.Error().Msgf("gracefulShutdown: %s", err.Error())
	}
}

func startApp(echo *echo.Echo, appCfg *infra.AppCfg) error {
	if err := di.Invoke(setRoute); err != nil {
		return err
	}

	return echo.StartServer(&http.Server{
		Addr:         appCfg.Address,
		ReadTimeout:  appCfg.ReadTimeout,
		WriteTimeout: appCfg.WriteTimeout,
	})
}

func gracefulShutdown(e *echo.Echo) {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	log.Info().Msg("shutting down server")

	if err := e.Shutdown(ctx); err != nil {
		log.Error().Msgf("echo shutdown: %s", err.Error())
	}

	log.Info().Msg("server gracefully stopped")
}
