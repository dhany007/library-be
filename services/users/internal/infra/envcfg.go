package infra

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func LoadAppCfg() (*AppCfg, error) {
	var cfg AppCfg
	prefix := "USERS_APP"
	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	return &cfg, nil
}

func LoadDatabaseCfg() (*DatabaseCfg, error) {
	var cfg DatabaseCfg
	prefix := "USERS_PG"
	if err := envconfig.Process(prefix, &cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", prefix, err)
	}

	return &cfg, nil
}
