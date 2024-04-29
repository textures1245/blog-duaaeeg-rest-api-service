package entities

type UsersPassport struct {
	Uuid      string `json:"uuid" db:"uuid"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
