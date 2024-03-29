package main

import (
	"fmt"

	"github.com/CRUD/api"
	"github.com/CRUD/config"
	"github.com/CRUD/pkg/logger"

	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api")

	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("new adapter error", logger.Error(err))

	}
	server := api.New(api.Option{
		Db:     connDB,
		Conf:   cfg,
		Logger: log,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
