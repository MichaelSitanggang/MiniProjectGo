package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type InputUseCase interface {
	Findall(userID int) ([]entities.Input_aktivitas, error)
	CreateAktip(userID int, akttivitas *entities.Input_aktivitas) error
}

type inputUseCase struct {
	repo repositories.AktivitasRepo
}

func NewInputUsecase(repo repositories.AktivitasRepo) InputUseCase {
	return inputUseCase{repo: repo}
}

func (iuc inputUseCase) Findall(userID int) ([]entities.Input_aktivitas, error) {
	return iuc.repo.GetAktivitasAll(userID)
}

func (iuc inputUseCase) CreateAktip(userID int, aktivitas *entities.Input_aktivitas) error {
	aktivitas.Total_jejak_karbon = aktivitas.Data_Aktivitas * float64(aktivitas.Konsumsi_energi_kwh)
	aktivitas.User_id = userID
	return iuc.repo.CreateAktivitas(aktivitas)
}
