package storage

import (
	"github.com/CRUD/pkg/logger"
	"github.com/CRUD/storage/postgres"
	"github.com/CRUD/storage/repo"
	"github.com/jackc/pgx/v4"
)

type IStorage interface {
	UserService() repo.UsersRepoInterface
	CountryService() repo.CountryRepoInterface
}

type storagePostgres struct {
	db             *pgx.Conn
	log            logger.Logger
	userService    repo.UsersRepoInterface
	countryService repo.CountryRepoInterface
}

func NewStoragePg(db *pgx.Conn, log logger.Logger) IStorage {
	return &storagePostgres{
		db:             db,
		log:            log,
		userService:    postgres.NewUsersRepasitory(db),
		countryService: postgres.NewCountryRepasitory(db),
	}
}

func (s *storagePostgres) UserService() repo.UsersRepoInterface {
	return s.userService
}

func (s *storagePostgres) CountryService() repo.CountryRepoInterface {
	return s.countryService
}
