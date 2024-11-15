package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/MiniProjectGo/services"

	"github.com/gin-gonic/gin"
)

// ini fitur ai
type ChatCotnroller struct {
	chatUseCase services.ChatUseCase
}

func NewControllerChat(chatUseCase services.ChatUseCase) *ChatCotnroller {
	return &ChatCotnroller{chatUseCase: chatUseCase}
}

func (cc *ChatCotnroller) ChatController(c *gin.Context) {
	var input struct {
		UserInput string `json:"user_input"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Input Data Invalid"})
		return
	}
	chat, err := cc.chatUseCase.ProsesChat(input.UserInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  "Gagal",
			"Kondisi": false,
			"Message": "Server eror",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Message": chat})
}

func (cc *ChatCotnroller) GetAllChats(c *gin.Context) {
	chats, err := cc.chatUseCase.GetAllChats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Message": chats})
}
