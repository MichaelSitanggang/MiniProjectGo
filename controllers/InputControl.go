package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

type InputControl struct {
	inputUse services.InputUseCase
}

func NewInputController(uc services.InputUseCase) *InputControl {
	return &InputControl{uc}
}

func (ic *InputControl) GetInputAktivitasAll(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := user_id.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	aktivitas, err := ic.inputUse.Findall(int(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aktivitas)
}

func (ic *InputControl) CreatedAktivitas(c *gin.Context) {
	// Ambil user_id dari context yang telah diset oleh middleware
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	userID, ok := user_id.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	userIDInt := int(userID)

	var input entities.Input_aktivitas
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.User_id = userIDInt

	if err := ic.inputUse.CreateAktip(userIDInt, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Aktivitas created successfully"})
}
