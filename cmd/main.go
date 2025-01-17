package main

import (
	"context"
	"fmt"

	"github.com/CRUD/api"
	"github.com/CRUD/config"
	"github.com/CRUD/pkg/logger"
	"github.com/CRUD/storage"
	"github.com/jackc/pgx/v4"

	gormadapter "github.com/casbin/gorm-adapter/v2"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	connDB, err := pgx.Connect(ctx, psqlString)
	// connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = gormadapter.NewAdapter("postgres", psqlString, true)
	if err != nil {
		log.Error("new adapter error", logger.Error(err))

	}

	storage := storage.NewStoragePg(connDB, log)

	server := api.New(api.Option{
		Db:      connDB,
		Conf:    cfg,
		Logger:  log,
		Storage: storage,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

}
