package main

import (
	"github.com/MichaelSitanggang/MiniProjectGo/config"
	"github.com/MichaelSitanggang/MiniProjectGo/controllers"
	middleware "github.com/MichaelSitanggang/MiniProjectGo/middlewares"
	"github.com/MichaelSitanggang/MiniProjectGo/repositories"
	"github.com/MichaelSitanggang/MiniProjectGo/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.CreateDatabase()
	//
	UserRepo := repositories.NewUserRepo(db)
	inputRepo := repositories.NewAktivitasRepo(db)
	//
	UserUseCase := services.NewUserUseCase(UserRepo)
	AktivitasUseCase := services.NewInputUsecase(inputRepo)
	//
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
		route.POST("", AktivitasController.CreatedAktivitas)
	}
	r.Run(":8000")
}

// Cek User All sudah pas atau belom
// dan minggu selanjut nya buat code unutk input aktivitas dlu
