package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/controllers"
	"github.com/nanda03dev/go2ms/services"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	userController := controllers.NewUserController(services.AppServices.User)
	orderController := controllers.NewOrderController(services.AppServices.Order)
	cityController := controllers.NewCityController(services.AppServices.City)
	productController := controllers.NewProductController(services.AppServices.Product)
	itemController := controllers.NewItemController(services.AppServices.Item)
	eventController := controllers.NewEventController(services.AppServices.Event)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.POST("/filter", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", orderController.CreateOrder)
		orderRoutes.POST("/filter", orderController.GetAllOrders)
		orderRoutes.GET("/:id", orderController.GetOrderByID)
		orderRoutes.PUT("/:id", orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", orderController.DeleteOrder)
	}

	cityRoutes := router.Group("/cities")
	{
		cityRoutes.POST("/", cityController.CreateCity)
		cityRoutes.POST("/filter", cityController.GetAllCities)
		cityRoutes.GET("/:id", cityController.GetCityByID)
		cityRoutes.PUT("/:id", cityController.UpdateCity)
		cityRoutes.DELETE("/:id", cityController.DeleteCity)
	}

	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.POST("/filter", productController.GetAllProducts)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	itemRoutes := router.Group("/items")
	{
		itemRoutes.POST("/", itemController.CreateItem)
		itemRoutes.POST("/filter", itemController.GetAllItems)
		itemRoutes.GET("/:id", itemController.GetItemByID)
		itemRoutes.PUT("/:id", itemController.UpdateItem)
		itemRoutes.DELETE("/:id", itemController.DeleteItem)
	}

	eventRoutes := router.Group("/events")
	{
		eventRoutes.POST("/", eventController.CreateEvent)
		eventRoutes.POST("/filter", eventController.GetAllEvents)
		eventRoutes.GET("/:id", eventController.GetEventByID)
		eventRoutes.PUT("/:id", eventController.UpdateEvent)
		eventRoutes.DELETE("/:id", eventController.DeleteEvent)
	}

	return router
}
