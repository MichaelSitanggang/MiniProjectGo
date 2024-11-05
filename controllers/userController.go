package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase services.UserUseCase
}

func NewController(usecase services.UserUseCase) *UserController {
	return &UserController{usecase}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Data tidak valid"})
		return
	}
	if err := uc.usecase.Regis(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Register Gagal"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Register Berhasil Lanjut Login"})
}
