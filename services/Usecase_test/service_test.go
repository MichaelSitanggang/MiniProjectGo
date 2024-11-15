package services

import (
	"testing"
	"time"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories/mocks"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAktip(t *testing.T) {
	mockRepo := new(mocks.AktivitasRepo)

	userID := 1
	inputAktivitas := &entities.Input_aktivitas{
		Data_Aktivitas:      10.0,
		Konsumsi_energi_kwh: 2.0,
	}

	mockRepo.On("CreateAktivitas", inputAktivitas).Return(nil)
	mockRepo.On("CreateHistory", mock.AnythingOfType("*entities.History")).Return(nil)

	usecase := services.NewInputUsecase(mockRepo)
	err := usecase.CreateAktip(userID, inputAktivitas)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindAllHistory(t *testing.T) {
	mockRepo := new(mocks.AktivitasRepo)

	mockRepo.On("GetAllHistoryByUserID", 1).Return([]entities.History{
		{User_id: 1, AktivitasID: 1, TotalKarbon: 100, CreatedAt: time.Now().String()},
	}, nil).Once()

	usecase := services.NewInputUsecase(mockRepo)
	history, err := usecase.FindAllHistory(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, history)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAktip(t *testing.T) {
	mockRepo := new(mocks.AktivitasRepo)
	mockRepo.On("DeleteAktivitas", 1, 1).Return(nil).Once()

	usecase := services.NewInputUsecase(mockRepo)
	err := usecase.DeleteAktip(1, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAktip(t *testing.T) {
	mocrepo := new(mocks.AktivitasRepo)

	mocrepo.On("FindbyId", 1, 1).Return(&entities.Input_aktivitas{
		Id: 1,
	}, nil).Once()
	mocrepo.On("UpdateAktivitas", mock.Anything).Return(nil).Once()
	mocrepo.On("CreateHistory", mock.Anything).Return(nil).Once()

	usecase := services.NewInputUsecase(mocrepo)
	aktivitas := &entities.Input_aktivitas{
		Data_Aktivitas:      10,
		Konsumsi_energi_kwh: 20,
	}

	err := usecase.UpdateAktip(1, 1, aktivitas)
	assert.NoError(t, err)
	mocrepo.AssertExpectations(t)
}

func TestGetAllAktivitas(t *testing.T) {
	mockRepo := new(mocks.AktivitasRepo)
	mockRepo.On("GetAktivitasAll", 1).Return([]entities.Input_aktivitas{
		{User_id: 1, Data_Aktivitas: 10.0, Konsumsi_energi_kwh: 2.0, Total_jejak_karbon: 20},
	}, nil).Once()
	usecase := services.NewInputUsecase(mockRepo)
	aktivitas, err := usecase.Findall(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, aktivitas)
	mockRepo.AssertExpectations(t)
}
