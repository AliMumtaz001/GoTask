package model

type User struct {
	ID       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
}

type UserCredentials struct {
	Email    string `json:"email" db:"email" binding:"required,email"`
    Password string `json:"password" binding:"required" db:"password"`
}