package services

import (
	"context"
	"fmt"
	"os"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ChatUseCase interface {
	CreateChat(userInput string) (entities.Chat, error)
	GetChat() ([]entities.Chat, error)
}

type chatUseCase struct {
	repo repositories.ChatRepo
}

func NewChatUseCase(repo repositories.ChatRepo) ChatUseCase {
	return &chatUseCase{repo: repo}
}

func (uc *chatUseCase) CreateChat(userInput string) (entities.Chat, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return entities.Chat{}, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userInput))
	if err != nil {
		return entities.Chat{}, err
	}

	if len(resp.Candidates) == 0 {
		return entities.Chat{}, err
	}
	aiRespon := ""
	for _, candidate := range resp.Candidates {
		if candidate.Content == nil {
			continue
		}
		for _, data := range candidate.Content.Parts {
			aiRespon += fmt.Sprintf("%s", data)
		}
	}
	chat := entities.Chat{
		UserInput: userInput,
		AiRespon:  aiRespon,
	}

	return chat, nil

}

func (uc *chatUseCase) GetChat() ([]entities.Chat, error) {
	return uc.repo.GetAllChat()
}
