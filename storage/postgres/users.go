package postgres

import (
	"github.com/CRUD/pkg/models"
	"github.com/CRUD/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type usersRepasitory struct {
	db *sqlx.DB
}

func NewUsersRepasitory(db *sqlx.DB) repo.UsersRepoInterface {
	return &usersRepasitory{
		db: db,
	}
}

func (r *usersRepasitory) Create(user *models.UserCreate) (*models.UserResponse, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	userResponse := models.UserResponse{}

	query := `INSERT INTO users(id,name,surname,patronymic,gender) VALUES ($1,$2,$3,$4,$5) RETURNING id,name,surname,patronymic,gender`
	err = r.db.QueryRow(query, id, user.Name, user.Surname, user.Patronymic, user.Gender).Scan(
		&userResponse.ID,
		&userResponse.Name,
		&userResponse.Surname,
		&userResponse.Patronymic,
		&userResponse.Gender,
	)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (r *usersRepasitory) Update(user *models.UpdateUser) (*models.UserResponse, error) {
	var (
		userResponse = models.UserResponse{}
		query        = `UPDATE users SET name=$1,surname=$2,patronymic=$3,gender=$4 WHERE id=$5 RETURNING id,name,surname,patronymic,gender`
	)

	err := r.db.QueryRow(query, user.Name, user.Surname, user.Patronymic, user.Gender, user.ID).Scan(
		&userResponse.ID,
		&userResponse.Name,
		&userResponse.Surname,
		&userResponse.Patronymic,
		&userResponse.Gender,
	)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (r *usersRepasitory) Get(id string) (*models.UserResponse, error) {
	var (
		userResponse = models.UserResponse{}
		query        = `SELECT id,name,surname,patronymic,gender FROM users WHERE id=$1`
	)

	err := r.db.QueryRow(query, id).Scan(
		&userResponse.ID,
		&userResponse.Name,
		&userResponse.Surname,
		&userResponse.Patronymic,
		&userResponse.Gender,
	)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (r *usersRepasitory) GetAll() (*[]models.UserResponse, error) {
	var (
		allUsers = []models.UserResponse{}
		query    = `SELECT id,name,surname,patronymic,gender from users`
	)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		singleUser := models.UserResponse{}
		err := rows.Scan(
			&singleUser.ID,
			&singleUser.Name,
			&singleUser.Surname,
			&singleUser.Patronymic,
			&singleUser.Gender,
		)

		if err != nil {
			return nil, err
		}
		allUsers = append(allUsers, singleUser)
	}

	return &allUsers, nil
}

func (r *usersRepasitory) DeleteUser(id string) error {
	var (
		query = `DELETE FROM users WHERE id=$1`
	)

	_, err := r.db.Exec(query, id)
	return err
}
