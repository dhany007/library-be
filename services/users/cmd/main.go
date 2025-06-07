package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	app "github.com/dhany007/library-be/services/users/internal"
	"github.com/dhany007/library-be/services/users/internal/handler/rest"
	"github.com/dhany007/library-be/services/users/internal/infra"
	"github.com/dhany007/library-be/services/users/pkg/di"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")

	}

	err = LoadApplicationConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationPackage()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationController()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	app.Start()
}

// LoadApplicationConfig load application config
func LoadApplicationConfig() error {
	err := di.Provide(infra.LoadAppCfg)
	if err != nil {
		return fmt.Errorf("LoadAppCfg: %s", err.Error())
	}

	return nil
}

// LoadApplicationPackage Load application package used by the application
func LoadApplicationPackage() error {
	err := di.Provide(infra.NewEcho)
	if err != nil {
		return fmt.Errorf("NewEcho: %s", err.Error())
	}

	return nil
}

// LoadApplicationController load application controller using uber dig
func LoadApplicationController() error {
	err := di.Provide(rest.NewHealthHandler)
	if err != nil {
		return fmt.Errorf("NewHealthHandler: %s", err.Error())
	}

	return nil
}
