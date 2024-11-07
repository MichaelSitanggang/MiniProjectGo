package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *entities.User) error
	GetByUsername(username string) (*entities.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db}
}

func (ur *userRepo) CreateUser(user *entities.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := ur.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
