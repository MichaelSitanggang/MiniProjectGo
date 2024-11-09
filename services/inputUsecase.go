package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type InputUseCase interface {
	Findall() (*entities.Input_aktivitas, error)
}

type inputUseCase struct {
	repo repositories.AktivitasRepo
}

func NewInputUsecase(repo repositories.AktivitasRepo) InputUseCase {
	return inputUseCase{repo: repo}
}

func (iuc inputUseCase) Findall() (*entities.Input_aktivitas, error) {
	return iuc.repo.GetAktivitasAll()
}
