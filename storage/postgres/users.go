package postgres

import (
	"context"

	"github.com/CRUD/pkg/models"
	"github.com/CRUD/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
)

type usersRepasitory struct {
	db *pgx.Conn
}

func NewUsersRepasitory(db *pgx.Conn) repo.UsersRepoInterface {
	return &usersRepasitory{
		db: db,
	}
}

func (r *usersRepasitory) Create(ctx context.Context, user *models.UserCreate) (*models.UserResponse, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	userResponse := models.UserResponse{}

	query := `INSERT INTO users(id,name,surname,patronymic,gender) VALUES ($1,$2,$3,$4,$5) RETURNING id,name,surname,patronymic,gender`
	err = r.db.QueryRow(ctx, query, id, user.Name, user.Surname, user.Patronymic, user.Gender).Scan(
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

func (r *usersRepasitory) Update(ctx context.Context, user *models.UpdateUser) (*models.UserResponse, error) {
	var (
		userResponse = models.UserResponse{}
		query        = `UPDATE users SET name=$1,surname=$2,patronymic=$3,gender=$4 WHERE id=$5 RETURNING id,name,surname,patronymic,gender`
	)

	err := r.db.QueryRow(ctx, query, user.Name, user.Surname, user.Patronymic, user.Gender, user.ID).Scan(
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

func (r *usersRepasitory) Get(ctx context.Context, id string) (*models.UserResponse, error) {
	var (
		userResponse = models.UserResponse{}
		query        = `SELECT id,name,surname,patronymic,gender FROM users WHERE id=$1`
	)

	err := r.db.QueryRow(ctx, query, id).Scan(
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

func (r *usersRepasitory) GetAll(ctx context.Context) (*[]models.UserResponse, error) {
	var (
		allUsers = []models.UserResponse{}
		query    = `SELECT id,name,surname,patronymic,gender from users`
	)

	rows, err := r.db.Query(ctx, query)
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

func (r *usersRepasitory) DeleteUser(ctx context.Context, id string) error {
	var (
		query = `DELETE FROM users WHERE id=$1`
	)

	_, err := r.db.Exec(ctx, query, id)
	return err
}
