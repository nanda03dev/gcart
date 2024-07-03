package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/oms/config"
	"github.com/nanda03dev/oms/controllers"
	"github.com/nanda03dev/oms/repositories"
	"github.com/nanda03dev/oms/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	return router
}
