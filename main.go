package main

import (
	"github.com/MichaelSitanggang/MiniProjectGo/config"
	"github.com/MichaelSitanggang/MiniProjectGo/controllers"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
	"github.com/MichaelSitanggang/MiniProjectGo/routes"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
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
	InputAktivitasController := controllers.NewInputController(InputAktivitasUseCase)
	ControlAktivitas := controllers.NewControlAktipitas(AktipitasUseCase)

	r := routes.SetupRouter(UserController, InputAktivitasController, ControlAktivitas, ChatController)

	r.Run(":8080")
}
