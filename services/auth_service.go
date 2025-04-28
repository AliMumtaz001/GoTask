package services

import (
	"errors"

	"github.com/AliMumtaz001/GoTask/api/repositories"
	model "github.com/AliMumtaz001/GoTask/models"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

// AuthenticateUser validate user cred.
func (service *AuthService) AuthenticateUser(credentials model.UserCredentials) (model.User, error) {
	user, err := service.UserRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		return user, err
	}

	if user.Password != credentials.Password {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}

// RegisterUser registers a new user.
func (service *AuthService) RegisterUser(user model.User) error {
	_, err := service.UserRepo.GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	if err.Error() != "user not found" {
		return err
	}

	return service.UserRepo.CreateUser(user)
}
