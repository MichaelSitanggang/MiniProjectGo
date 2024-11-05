package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type UserRepo interface {
	Register(user *entities.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (ur *userRepo) Register(user *entities.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
