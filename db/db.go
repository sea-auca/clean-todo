package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-rel/postgres"
	"github.com/go-rel/rel"
	_ "github.com/lib/pq"
	"github.com/sea-auca/clean-todo/config"
	"github.com/sea-auca/clean-todo/utils"
)

var ErrNoCredentials = errors.New("no credentials were supplied for production environment")

func ConnectRelDatabase(conf *config.AppConfig, shutdown *utils.Shutdown) (rel.Repository, error) {

	// read credentials from environment variables for production
	// in development this code will not be executed and will use
	// unprotected credentials from config.dev.yml
	if !conf.IsDevelopmentConfig {
		password, exist_pass := os.LookupEnv("POSTGRESQL_USERNAME")
		username, exist_user := os.LookupEnv("POSTGRESQL_PASSWORD")
		if !(exist_pass && exist_user) {
			return nil, ErrNoCredentials
		}
		conf.DatabaseConfig.User = username
		conf.DatabaseConfig.Password = password
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DatabaseConfig.User, conf.DatabaseConfig.Password, conf.DatabaseConfig.Host, conf.DatabaseConfig.Port, conf.DatabaseConfig.Database)

	adapter, err := postgres.Open(dsn)
	if err != nil {
		return nil, err
	}
	shutdown.Add(adapter.Close) // add close to the global shutdown
	repo := rel.New(adapter)

	return repo, nil
}
