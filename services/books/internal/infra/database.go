package infra

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type (
	Databases struct {
		dig.Out
		Pg *sql.DB
	}

	DatabaseCfgs struct {
		dig.In
		Pg *DatabaseCfg
	}

	DatabaseCfg struct {
		DBName string `envconfig:"DBNAME" required:"true" default:"postgres"`
		DBUser string `envconfig:"DBUSER" required:"true" default:"postgres"`
		DBPass string `envconfig:"DBPASS" required:"true" default:"postgres"`
		Host   string `envconfig:"HOST" required:"true" default:"localhost"`
		Port   string `envconfig:"PORT" required:"true" default:"5432"`

		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"20" required:"true"`
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"5" required:"true"`
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"15m" required:"true"`
	}
)

func NewDatabases(cfgs DatabaseCfgs) Databases {
	return Databases{
		Pg: openPostgres(cfgs.Pg),
	}
}

func openPostgres(p *DatabaseCfg) *sql.DB {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.DBUser, url.QueryEscape(p.DBPass), p.Host, p.Port, p.DBName,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal().Err(err).Msg("postgres: open")
	}

	db.SetConnMaxLifetime(p.ConnMaxLifetime)
	db.SetMaxIdleConns(p.MaxIdleConns)
	db.SetMaxOpenConns(p.MaxOpenConns)

	if err = db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("postgres: ping")
	}

	return db
}
