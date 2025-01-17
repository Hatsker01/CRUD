package v1

import (
	"github.com/CRUD/config"
	"github.com/CRUD/pkg/logger"
	"github.com/CRUD/storage"
	"github.com/jackc/pgx/v4"
)

type handlerV1 struct {
	db      *pgx.Conn
	log     logger.Logger
	cfg     config.Config
	storage storage.IStorage
}

type HandlerV1Config struct {
	Db      *pgx.Conn
	Logger  logger.Logger
	Cfg     config.Config
	Storage storage.IStorage
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		db:      c.Db,
		log:     c.Logger,
		cfg:     c.Cfg,
		storage: c.Storage,
	}
}
