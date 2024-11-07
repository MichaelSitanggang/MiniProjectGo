package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type UserUseCase interface {
	RegisterUser(user *entities.User) error
	Login(username, password string) (*entities.User, error)
}

type userUseCase struct {
	repo repositories.UserRepo
}

func NewUserUseCase(repo repositories.UserRepo) UserUseCase {
	return &userUseCase{repo}
}

func (uc userUseCase) RegisterUser(user *entities.User) error {
	return uc.repo.CreateUser(user)
}

func (uc userUseCase) Login(username, password string) (*entities.User, error) {
	user, err := uc.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, err
	}
	return user, nil
}
