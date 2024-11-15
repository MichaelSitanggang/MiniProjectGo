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
	//
	ChatRepo := repositories.NewChatRepo(db)
	UserRepo := repositories.NewUserRepo(db)
	inputRepo := repositories.NewAktivitasRepo(db)
	//
	ChatUseCase := services.NewUseCaseChat(ChatRepo)
	UserUseCase := services.NewUserUseCase(UserRepo)
	InputAktivitasUseCase := services.NewInputUsecase(inputRepo)
	AktipitasUseCase := services.NewUseCaseAktivitas(repositories.NewRepoAktivitas(db))
	//
	ChatController := controllers.NewControllerChat(ChatUseCase)
	UserController := controllers.NewController(UserUseCase)
	AktivitasController := controllers.NewInputController(AktivitasUseCase)
	//
	r := gin.Default()
	r.POST("/register", UserController.RegisterUser)
	r.POST("/login", UserController.LoginUser)

	route := r.Group("/inputAktivitas")
	route.Use(middleware.AuthMiddleware())
	{
		route.GET("", AktivitasController.GetInputAktivitasAll)
		route.POST("inputAktivitas", AktivitasController.CreatedAktivitas)
	}
	r.Run(":8080")
}
