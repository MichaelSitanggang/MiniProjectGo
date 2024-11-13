package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type RepoAktivitas interface {
	GetAllAktivitas() ([]*entities.Aktivitas, error)
}

type repoAktivitas struct {
	db *gorm.DB
}

func NewRepoAktivitas(db *gorm.DB) RepoAktivitas {
	return &repoAktivitas{db: db}
}

func (ra repoAktivitas) GetAllAktivitas() ([]*entities.Aktivitas, error) {
	var aktivitas []*entities.Aktivitas
	err := ra.db.Find(&aktivitas).Error
	if err != nil {
		return nil, err
	}
	return aktivitas, nil
}
