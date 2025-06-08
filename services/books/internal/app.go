package internal

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/dhany007/library-be/proto/bookpb"
	grpcHandler "github.com/dhany007/library-be/services/books/internal/handler/grpc"
	"github.com/dhany007/library-be/services/books/internal/infra"
	"github.com/dhany007/library-be/services/books/internal/services"
	"github.com/dhany007/library-be/services/books/pkg/di"
)

var exitSigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

func Start() {
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, exitSigs...)

	go func() {
		defer func() { exitCh <- syscall.SIGTERM }()
		if err := di.Invoke(func(e *echo.Echo, cfg *infra.AppCfg, bookSvc services.BookService) error {
			return startApp(e, cfg, bookSvc)
		}); err != nil {
			log.Fatal().Msgf("startApp: %s", err.Error())
		}
	}()
	<-exitCh

	if err := di.Invoke(gracefulShutdown); err != nil {
		log.Error().Msgf("gracefulShutdown: %s", err.Error())
	}
}

func startApp(echo *echo.Echo, appCfg *infra.AppCfg, bookSvc services.BookService) error {
	go func() {
		startGRPCServer(bookSvc)
	}()

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

func startGRPCServer(svc services.BookService) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen on :50052")
	}

	grpcServer := grpc.NewServer()
	bookpb.RegisterBookServiceServer(grpcServer, grpcHandler.NewBookGRPCHandler(svc))

	log.Info().Msg("gRPC server running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("gRPC server error")
	}
}
