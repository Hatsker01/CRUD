package models

type CreateCountry struct {
	UserID      string  `json:"user_id"`
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type CountryResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type UpdateCountry struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type Countries struct {
	Countries []CountryResponse `json:"countries"`
}

type UserCountry struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
