package repo

import "github.com/CRUD/pkg/models"

type UsersRepoInterface interface {
	Create(user *models.UserCreate) (*models.UserResponse, error)
	Update(user *models.UpdateUser) (*models.UserResponse, error)
	Get(id string) (*models.UserResponse, error)
	GetAll() (*[]models.UserResponse, error)
	DeleteUser(id string) error
}

type CountryRepoInterface interface {
	Create(country *models.CreateCountry) (*models.CountryResponse, error)
	Get(id string) (*models.CountryResponse, error)
	Update(country *models.UpdateCountry) (*models.CountryResponse, error)
	Delete(id string) error
	GetUserCountry(userID string) (*[]models.CountryResponse, error)
	GetUserWithCountry(userID string) (*[]models.UserCountry, error)
}
