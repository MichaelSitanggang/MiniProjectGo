package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type AktivitasRepo interface {
	GetAktivitasAll() (*entities.Input_aktivitas, error)
}

type aktivitasRepo struct {
	db *gorm.DB
}

func NewAktivitasRepo(db *gorm.DB) aktivitasRepo {
	return aktivitasRepo{db: db}
}

func (ar aktivitasRepo) GetAktivitasAll() (*entities.Input_aktivitas, error) {
	var aktivitas *entities.Input_aktivitas
	ar.db.Find(&aktivitas)
	return aktivitas, nil
}
