package postgres

import (
	"github.com/CRUD/pkg/models"
	"github.com/CRUD/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type countryRepasitory struct {
	db *sqlx.DB
}

func NewCountryRepasitory(db *sqlx.DB) repo.CountryRepoInterface {
	return &countryRepasitory{
		db: db,
	}
}

func (r *countryRepasitory) Create(country *models.CreateCountry) (*models.CountryResponse, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	var (
		query      = `INSERT INTO nationalize(id,user_id,country_id,probability) VALUES ($1,$2,$3,$4) RETURNING id,user_id,country_id,probability`
		newCountry models.CountryResponse
	)
	err = r.db.QueryRow(query, id, country.UserID, country.CountryID, country.Probability).Scan(
		&newCountry.ID,
		&newCountry.UserID,
		&newCountry.CountryID,
		&newCountry.Probability,
	)

	if err != nil {
		return nil, err
	}

	return &newCountry, nil
}

func (r *countryRepasitory) Get(id string) (*models.CountryResponse, error) {
	var (
		query   = `SELECT id,user_id,country_id,probability from nationalize where id=$1`
		country models.CountryResponse
	)

	err := r.db.QueryRow(query, id).Scan(
		&country.ID,
		&country.UserID,
		&country.CountryID,
		&country.Probability,
	)
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (r *countryRepasitory) Update(country *models.UpdateCountry) (*models.CountryResponse, error) {
	var (
		query          = `UPDATE nationalize SET user_id=$1,country_id=$2,probability=$3 WHERE id=$4 RETURNING id,user_id,country_id,probability`
		updatedCountry models.CountryResponse
	)

	err := r.db.QueryRow(query, country.UserID, country.CountryID, country.Probability, country.ID).Scan(
		&updatedCountry.ID,
		&updatedCountry.UserID,
		&updatedCountry.CountryID,
		&updatedCountry.Probability,
	)

	if err != nil {
		return nil, err
	}

	return &updatedCountry, nil
}

func (r *countryRepasitory) Delete(id string) error {
	query := `DELETE FROM nationalize WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *countryRepasitory) GetUserCountry(userID string) (*[]models.CountryResponse, error) {
	var (
		query     = `SELECT id,user_id,country_id,probability from nationalize where user_id=$1`
		countries []models.CountryResponse
	)

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		country := models.CountryResponse{}
		if err := rows.Scan(
			&country.ID,
			&country.UserID,
			&country.CountryID,
			&country.Probability,
		); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	return &countries, nil
}

func (r *countryRepasitory) GetUserWithCountry(userID string) (*[]models.UserCountry, error) {
	var (
		query     = `SELECT country_id,probability from nationalize where user_id=$1`
		countries []models.UserCountry
	)

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		country := models.UserCountry{}
		if err := rows.Scan(
			&country.CountryID,
			&country.Probability,
		); err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}

	return &countries, nil
}
