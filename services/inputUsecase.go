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

// GetAll(page, limit int) ([]entities.Category, *entities.Pagination, error)

// func (s *categoryService) GetAll(page, limit int) ([]entities.Category, *entities.Pagination, error) {
//     var categories []entities.Category
//     totalItems, err := s.categoryRepo.Count()
//     if err != nil {
//         return nil, nil, err
//     }
//     if err := s.categoryRepo.FindAll(&categories, page, limit); err != nil {
//         return nil, nil, err
//     }
//     totalPages := int((totalItems + int64(limit) - 1) / int64(limit))
//     pagination := &entities.Pagination{
//         CurrentPage: page,
//         TotalPages:  totalPages,
//         TotalItems:  totalItems,
//     }
//     return categories, pagination, nil
// }

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
	active.Data_Aktivitas = aktivitas.Data_Aktivitas
	active.Konsumsi_energi_kwh = aktivitas.Konsumsi_energi_kwh
	active.Total_jejak_karbon = aktivitas.Data_Aktivitas * float64(aktivitas.Konsumsi_energi_kwh)

	if err := iuc.repo.UpdateAktivitas(active); err != nil {
		return err
	}
	history := entities.History{
		User_id:     userID,
		AktivitasID: active.Id,
		TotalKarbon: active.Total_jejak_karbon,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	return iuc.repo.CreateHistory(&history)
}

func (iuc inputUseCase) DeleteAktip(id int, userID int) error {
	return iuc.repo.DeleteAktivitas(id, userID)
}

func (iuc inputUseCase) FindAllHistory(userID int) ([]entities.History, error) {
	return iuc.repo.GetAllHistoryByUserID(userID)
}
