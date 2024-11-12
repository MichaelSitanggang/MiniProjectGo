package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type AktivitasRepo interface {
	GetAktivitasAll(userID int) ([]entities.Input_aktivitas, error)
	CreateAktivitas(aktivitas *entities.Input_aktivitas) error
}

type aktivitasRepo struct {
	db *gorm.DB
}

func NewAktivitasRepo(db *gorm.DB) aktivitasRepo {
	return aktivitasRepo{db: db}
}

func (ar aktivitasRepo) GetAktivitasAll(userID int) ([]entities.Input_aktivitas, error) {
	var aktivitas []entities.Input_aktivitas
	if err := ar.db.Where("user_id = ?", userID).Find(&aktivitas).Error; err != nil {
		return nil, err
	}
	return aktivitas, nil
}

func (ar aktivitasRepo) CreateAktivitas(aktivitas *entities.Input_aktivitas) error {
	return ar.db.Create(aktivitas).Error
}
