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
	ProsesChat(userInput string) (entities.Chat, error)
	GetAllChats() ([]entities.Chat, error)
}

type chatUseCase struct {
	repoC repositories.ChatRepo
}

// ini adalah fitur ai
func NewUseCaseChat(repoC repositories.ChatRepo) ChatUseCase {
	return &chatUseCase{repoC: repoC}
}

func (cuc *chatUseCase) ProsesChat(userInput string) (entities.Chat, error) {
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
	aiResponse := ""
	for _, candidate := range resp.Candidates {
		if candidate.Content == nil {
			continue
		}
		for _, part := range candidate.Content.Parts {
			aiResponse += fmt.Sprintf("%v", part)
		}
	}
	chat := entities.Chat{
		UserInput: userInput,
		AiRespon:  aiResponse,
	}
	if err := cuc.repoC.SaveChat(chat); err != nil {
		return entities.Chat{}, err
	}

	return chat, nil

}

func (cuc *chatUseCase) GetAllChats() ([]entities.Chat, error) {
	return cuc.repoC.GetAllChat()
}
