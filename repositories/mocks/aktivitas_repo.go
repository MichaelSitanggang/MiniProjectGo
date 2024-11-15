package mocks

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/stretchr/testify/mock"
)

type AktivitasRepo struct {
	mock.Mock
}

func (m *AktivitasRepo) GetAktivitasAll(userID int) ([]entities.Input_aktivitas, error) {
	test := m.Called(userID)
	return test.Get(0).([]entities.Input_aktivitas), test.Error(1)
}

func (m *AktivitasRepo) FindbyId(id, userID int) (*entities.Input_aktivitas, error) {
	args := m.Called(id, userID)
	return args.Get(0).(*entities.Input_aktivitas), args.Error(1)
}

func (m *AktivitasRepo) CreateAktivitas(aktivitas *entities.Input_aktivitas) error {
	test := m.Called(aktivitas)
	return test.Error(0)
}

func (m *AktivitasRepo) UpdateAktivitas(aktivitas *entities.Input_aktivitas) error {
	test := m.Called(aktivitas)
	return test.Error(0)
}

func (m *AktivitasRepo) DeleteAktivitas(id int, userID int) error {
	test := m.Called(id, userID)
	return test.Error(0)
}

func (m *AktivitasRepo) CreateHistory(history *entities.History) error {
	test := m.Called(history)
	return test.Error(0)
}

func (m *AktivitasRepo) GetAllHistoryByUserID(userID int) ([]entities.History, error) {
	test := m.Called(userID)
	return test.Get(0).([]entities.History), test.Error(1)
}
