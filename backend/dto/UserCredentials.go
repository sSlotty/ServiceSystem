package dto

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterCredentials struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ConfPassword string `json:"confPassword"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
}
