package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type InputUseCase interface {
	Findall(userID int) ([]entities.Input_aktivitas, error)
	CreateAktip(userID int, akttivitas *entities.Input_aktivitas) error
	UpdateAktip(id int, userID int, aktivitas *entities.Input_aktivitas) error
	DeleteAktip(id int, userID int) error
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

func (iuc inputUseCase) UpdateAktip(id int, userID int, aktivitas *entities.Input_aktivitas) error {
	active, err := iuc.repo.FindbyId(id, userID)
	if err != nil {
		return err
	}
	active.Data_Aktivitas = aktivitas.Data_Aktivitas
	active.Konsumsi_energi_kwh = aktivitas.Konsumsi_energi_kwh
	active.Total_jejak_karbon = aktivitas.Data_Aktivitas * float64(aktivitas.Konsumsi_energi_kwh)

	return iuc.repo.UpdateAktivitas(active)
}

func (iuc inputUseCase) DeleteAktip(id int, userID int) error {
	return iuc.repo.DeleteAktivitas(id, userID)
}
