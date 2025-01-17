package repo

import (
	"context"

	"github.com/CRUD/pkg/models"
)

type UsersRepoInterface interface {
	Create(ctx context.Context, user *models.UserCreate) (*models.UserResponse, error)
	Update(ctx context.Context, user *models.UpdateUser) (*models.UserResponse, error)
	Get(ctx context.Context, id string) (*models.UserResponse, error)
	GetAll(ctx context.Context) (*[]models.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
}

type CountryRepoInterface interface {
	Create(ctx context.Context, country *models.CreateCountry) (*models.CountryResponse, error)
	Get(ctx context.Context, id string) (*models.CountryResponse, error)
	Update(ctx context.Context, country *models.UpdateCountry) (*models.CountryResponse, error)
	Delete(ctx context.Context, id string) error
	GetUserCountry(ctx context.Context, userID string) (*[]models.CountryResponse, error)
	GetUserWithCountry(ctx context.Context, userID string) (*[]models.UserCountry, error)
}
