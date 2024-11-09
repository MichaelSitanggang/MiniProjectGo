package controllers

import (
	"net/http"

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
	input, err := ic.inputUse.Findall()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message ": "Data tidak ada"})
	}
	c.JSON(http.StatusOK, input)
}
