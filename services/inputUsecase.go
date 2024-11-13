package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type InputUseCase interface {
	Findall(userID int) ([]entities.Input_aktivitas, error)
	CreateAktip(userID int, akttivitas *entities.Input_aktivitas) error
	Update(userID int, aktivitasID int, aktivitas *entities.Input_aktivitas) error
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

func (iuc inputUseCase) Update(aktivitasID int, userID int, aktivitas *entities.Input_aktivitas) error {
	activ, err := iuc.repo.FindByID(aktivitasID, userID)
	if err != nil {
		return err
	}
	activ.Data_Aktivitas = aktivitas.Data_Aktivitas
	activ.Konsumsi_energi_kwh = aktivitas.Konsumsi_energi_kwh
	activ.Total_jejak_karbon = aktivitas.Data_Aktivitas * float64(aktivitas.Konsumsi_energi_kwh)

	return iuc.repo.UpdateAktivitas(activ)
}
