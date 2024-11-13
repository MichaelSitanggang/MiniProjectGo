package routes

import (
	"github.com/MichaelSitanggang/MiniProjectGo/controllers"
	middleware "github.com/MichaelSitanggang/MiniProjectGo/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, inputAktivitasController *controllers.InputControl, aktivitasController *controllers.ControllerAktip) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.LoginUser)

	activityRoutes := r.Group("/activity")
	activityRoutes.Use(middleware.AuthMiddleware())
	{
		activityRoutes.GET("", inputAktivitasController.GetInputAktivitasAll)
		activityRoutes.POST("", inputAktivitasController.CreatedAktivitas)
		activityRoutes.PUT("/:id", inputAktivitasController.UpdatedAktivitas)
		activityRoutes.DELETE("/:id", inputAktivitasController.DeletedAktivitas)
	}

	historyRoutes := r.Group("/history")
	historyRoutes.Use(middleware.AuthMiddleware())
	{
		historyRoutes.GET("", inputAktivitasController.GetHistoryAll)
	}

	typeActivityRoutes := r.Group("/typeactivity")
	typeActivityRoutes.Use(middleware.AuthMiddleware())
	{
		typeActivityRoutes.GET("", aktivitasController.GetAllAktip)
	}

	return r
}
