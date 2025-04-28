package repositories

import (
	"database/sql"
	"errors"
	"github.com/AliMumtaz001/GoTask/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := `SELECT email, password FROM employee WHERE email = $1`
	err := repo.DB.QueryRow(query, email).Scan(&user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user model.User) error {
	query := `INSERT INTO employee (email, password) VALUES ($1, $2)`
	_, err := repo.DB.Exec(query, user.Email, user.Password)
	return err
}

