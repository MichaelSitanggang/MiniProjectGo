package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	usecase services.ChatUseCase
}

func NewChatController(usecase services.ChatUseCase) *ChatController {
	return &ChatController{usecase: usecase}
}

func (u *ChatController) CreatedChat(c *gin.Context) {
	var inputChat struct {
		UserInput string `json:"user_input"`
	}
	if err := c.ShouldBindJSON(&inputChat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Input Data Invalid"})
		return
	}
	chats, err := u.usecase.CreateChat(inputChat.UserInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Eror pada server"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Message": chats})
}

func (u *ChatController) GetChatAlls(c *gin.Context) {
	chats, err := u.usecase.GetChat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "Gagal", "Kondisi": false, "Message": "Eror pada server"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "Berhasil", "Kondisi": true, "Message": chats})
}
