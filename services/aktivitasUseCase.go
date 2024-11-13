package services

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
)

type UseCaseAktivitas interface {
	AllAktivitas() ([]*entities.Aktivitas, error)
}

type useCaseAktivitas struct {
	ur repositories.RepoAktivitas
}

func NewUseCaseAktivitas(ur repositories.RepoAktivitas) UseCaseAktivitas {
	return &useCaseAktivitas{ur: ur}
}

func (uk useCaseAktivitas) AllAktivitas() ([]*entities.Aktivitas, error) {
	return uk.ur.GetAllAktivitas()
}
