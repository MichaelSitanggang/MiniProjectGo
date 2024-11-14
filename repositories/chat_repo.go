package repositories

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/gorm"
)

type ChatRepo interface {
	SaveChat(chat entities.Chat) error
	GetAllChat() ([]entities.Chat, error)
}

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) ChatRepo {
	return &chatRepo{db: db}
}

func (r *chatRepo) SaveChat(chat entities.Chat) error {
	return r.db.Create(&chat).Error
}

func (r *chatRepo) GetAllChat() ([]entities.Chat, error) {
	var chats []entities.Chat
	err := r.db.Find(chats).Error
	return chats, err
}
