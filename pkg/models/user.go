package models

type UserCreate struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Gender     string `json:"gender"`
}

type UserResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Gender     string `json:"gender"`
}

type Gender struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Gender     string `json:"gender"`
}

type Users struct {
	Users []UserResponse `json:"users"`
}

type UserWithCountry struct {
	Name    string        `json:"name"`
	Country []UserCountry `json:"country"`
}

type Empty struct{}
