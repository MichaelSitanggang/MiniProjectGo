package services

import (
	"time"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type InputUseCase interface {
	Findall(userID int) ([]entities.Input_aktivitas, error)
	CreateAktip(userID int, akttivitas *entities.Input_aktivitas) error
	UpdateAktip(id int, userID int, aktivitas *entities.Input_aktivitas) error
	DeleteAktip(id int, userID int) error
	FindAllHistory(userID int) ([]entities.History, error)
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
	// Hitung Total Jejak Karbon
	aktivitas.Total_jejak_karbon = aktivitas.Data_Aktivitas * float64(aktivitas.Konsumsi_energi_kwh)
	aktivitas.User_id = userID

	// Simpan Aktivitas
	if err := iuc.repo.CreateAktivitas(aktivitas); err != nil {
		return err
	}

	// Tambahkan ke History
	history := entities.History{
		User_id:     userID,
		AktivitasID: aktivitas.Id,
		TotalKarbon: aktivitas.Total_jejak_karbon,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
	if err := iuc.repo.CreateHistory(&history); err != nil {
		return err
	}

	return nil
}

//teersolve

func (iuc inputUseCase) UpdateAktip(id int, userID int, aktivitas *entities.Input_aktivitas) error {
	// Temukan aktivitas yang ingin diupdate berdasarkan ID dan userID
	active, err := iuc.repo.FindbyId(id, userID)
	if err != nil {
		return err
	}
	// Update aktivitas yang ada dengan data baru
	active.Data_Aktivitas = aktivitas.Data_Aktivitas
	active.Konsumsi_energi_kwh = aktivitas.Konsumsi_energi_kwh
	active.Total_jejak_karbon = aktivitas.Data_Aktivitas * float64(aktivitas.Konsumsi_energi_kwh)

	// Lakukan update aktivitas di database
	if err := iuc.repo.UpdateAktivitas(active); err != nil {
		return err
	}

	// Buat history baru setelah aktivitas di-update
	history := entities.History{
		User_id:     userID,
		AktivitasID: active.Id,
		TotalKarbon: active.Total_jejak_karbon,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339), // Jika diperlukan
	}

	return iuc.repo.CreateHistory(&history)
}

func (iuc inputUseCase) DeleteAktip(id int, userID int) error {
	return iuc.repo.DeleteAktivitas(id, userID)
}

func (iuc inputUseCase) FindAllHistory(userID int) ([]entities.History, error) {
	return iuc.repo.GetAllHistoryByUserID(userID)
}
