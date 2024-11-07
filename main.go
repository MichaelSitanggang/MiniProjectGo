package main

import (
	"github.com/MichaelSitanggang/MiniProjectGo/config"
	"github.com/MichaelSitanggang/MiniProjectGo/controllers"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.CreateDatabase()
	UserRepo := repositories.NewUserRepo(db)
	UserUseCase := services.NewUserUseCase(UserRepo)
	UserController := controllers.NewController(UserUseCase)
	r := gin.Default()
	r.POST("/register", UserController.RegisterUser)
	r.POST("/login", UserController.LoginUser)
	r.Run(":8000")
}
