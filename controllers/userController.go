package controllers

import (
	"net/http"

	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	middleware "github.com/MichaelSitanggang/MiniProjectGo/middlewares"
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
	if err := uc.usecase.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Register Gagal"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Message": "Register Berhasil Lanjut Login"})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	// var user entities.User
	var credential struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindBodyWithJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := uc.usecase.Login(credential.Username, credential.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := middleware.GenerateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menggunakan token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
