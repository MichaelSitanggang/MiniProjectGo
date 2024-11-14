package repositories

import "github.com/MichaelSitanggang/MiniProjectGo/entities"

type ChatRepo interface {
	SaveChat(chat entities.Chat) error
}
