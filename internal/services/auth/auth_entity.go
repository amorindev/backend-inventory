package auth

type UserEntity struct {
	Email   string  `json:"email" db:"user_email"`
	Password string `json:"password" db:"user_pass"`
}