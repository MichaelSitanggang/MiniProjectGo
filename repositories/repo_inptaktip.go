package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type AktivitasRepo interface {
	GetAktivitasAll(userID int) ([]entities.Input_aktivitas, error)
	FindbyId(id, userID int) (*entities.Input_aktivitas, error)
	CreateAktivitas(aktivitas *entities.Input_aktivitas) error
	UpdateAktivitas(aktivitas *entities.Input_aktivitas) error
	DeleteAktivitas(id int, userID int) error
<<<<<<< HEAD
	CreateHistory(history *entities.History) error
	GetAllHistoryByUserID(userID int) ([]entities.History, error)
=======
>>>>>>> fiturAktivitas
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

func (ar aktivitasRepo) FindbyId(id, userID int) (*entities.Input_aktivitas, error) {
	var aktivitas entities.Input_aktivitas
	if err := ar.db.Where("id = ? AND user_id = ?", id, userID).First(&aktivitas).Error; err != nil {
		return nil, err
	}
	return &aktivitas, nil
}
func (ar aktivitasRepo) CreateAktivitas(aktivitas *entities.Input_aktivitas) error {
	return ar.db.Create(aktivitas).Error
}

func (ar aktivitasRepo) UpdateAktivitas(aktivitas *entities.Input_aktivitas) error {
	return ar.db.Save(aktivitas).Error
}

func (ar aktivitasRepo) DeleteAktivitas(id int, userID int) error {
	var aktivitas entities.Input_aktivitas
	return ar.db.Where("id = ? AND user_id = ?", id, userID).Delete(&aktivitas).Error
}
<<<<<<< HEAD

func (ar aktivitasRepo) CreateHistory(history *entities.History) error {
	return ar.db.Create(history).Error
}

func (ar aktivitasRepo) GetAllHistoryByUserID(userID int) ([]entities.History, error) {
	var histories []entities.History
	err := ar.db.Where("user_id = ?", userID).Find(&histories).Error
	if err != nil {
		return nil, err
	}
	return histories, nil
}
=======
>>>>>>> fiturAktivitas
