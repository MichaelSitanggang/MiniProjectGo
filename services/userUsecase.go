package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type UserUseCase interface {
	Regis(user *entities.User) error
}

type userUseCase struct {
	repo repositories.UserRepo
}

func NewUserUseCase(repo repositories.UserRepo) UserUseCase {
	return &userUseCase{repo}
}

func (uc userUseCase) Regis(user *entities.User) error {
	return uc.repo.Register(user)
}
