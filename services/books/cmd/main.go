package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	app "github.com/dhany007/library-be/services/books/internal"
	"github.com/dhany007/library-be/services/books/internal/handler/rest"
	"github.com/dhany007/library-be/services/books/internal/infra"
	"github.com/dhany007/library-be/services/books/internal/repository/postgres"
	"github.com/dhany007/library-be/services/books/internal/services"
	"github.com/dhany007/library-be/services/books/pkg/di"
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

	err = LoadApplicationRepository()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = LoadApplicationService()
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

	err = di.Provide(infra.LoadDatabaseCfg)
	if err != nil {
		return fmt.Errorf("LoadDatabaseCfg: %s", err.Error())
	}
	return nil
}

// LoadApplicationPackage Load application package used by the application
func LoadApplicationPackage() error {
	err := di.Provide(infra.NewEcho)
	if err != nil {
		return fmt.Errorf("NewEcho: %s", err.Error())
	}

	err = di.Provide(infra.NewDatabases)
	if err != nil {
		return fmt.Errorf("NewDatabases: %s", err.Error())
	}

	return nil
}

// LoadApplicationRepository load repository using ubed dig
func LoadApplicationRepository() error {
	err := di.Provide(postgres.NewAuthorRepository)
	if err != nil {
		return fmt.Errorf("NewAuthorRepository: %s", err.Error())
	}

	err = di.Provide(postgres.NewCategoryRepository)
	if err != nil {
		return fmt.Errorf("NewCategoryRepository: %s", err.Error())
	}

	return nil
}

// LoadApplicationService Load service or usecase of the application using uber dig
func LoadApplicationService() error {
	err := di.Provide(services.NewAuthorService)
	if err != nil {
		return fmt.Errorf("NewAuthorService: %s", err.Error())
	}

	err = di.Provide(services.NewCategoryService)
	if err != nil {
		return fmt.Errorf("NewCategoryService: %s", err.Error())
	}

	return nil
}

// LoadApplicationController load application controller using uber dig
func LoadApplicationController() error {
	err := di.Provide(rest.NewHealthHandler)
	if err != nil {
		return fmt.Errorf("NewHealthHandler: %s", err.Error())
	}

	err = di.Provide(rest.NewAuthorHandler)
	if err != nil {
		return fmt.Errorf("NewAuthorHandler: %s", err.Error())
	}

	err = di.Provide(rest.NewCategoryHandler)
	if err != nil {
		return fmt.Errorf("NewCategoryHandler: %s", err.Error())
	}

	return nil
}
