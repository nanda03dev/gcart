package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nanda03dev/go2ms/config"
	"github.com/nanda03dev/go2ms/controllers"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userRepository := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	orderRepository := repositories.NewOrderRepository(config.DB)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	cityRepository := repositories.NewCityRepository(config.DB)
	cityService := services.NewCityService(cityRepository)
	cityController := controllers.NewCityController(cityService)

	productRepository := repositories.NewProductRepository(config.DB)
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	itemRepository := repositories.NewItemRepository(config.DB)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	orderRoutes := router.Group("/orders")
	{
		orderRoutes.POST("/", orderController.CreateOrder)
		orderRoutes.GET("/", orderController.GetAllOrders)
		orderRoutes.GET("/:id", orderController.GetOrderByID)
		orderRoutes.PUT("/:id", orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", orderController.DeleteOrder)
	}

	cityRoutes := router.Group("/cities")
	{
		cityRoutes.POST("/", cityController.CreateCity)
		cityRoutes.GET("/", cityController.GetAllCities)
		cityRoutes.GET("/:id", cityController.GetCityByID)
		cityRoutes.PUT("/:id", cityController.UpdateCity)
		cityRoutes.DELETE("/:id", cityController.DeleteCity)
	}

	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", productController.CreateProduct)
		productRoutes.GET("/", productController.GetAllProducts)
		productRoutes.GET("/:id", productController.GetProductByID)
		productRoutes.PUT("/:id", productController.UpdateProduct)
		productRoutes.DELETE("/:id", productController.DeleteProduct)
	}

	itemRoutes := router.Group("/items")
	{
		itemRoutes.POST("/", itemController.CreateItem)
		itemRoutes.GET("/", itemController.GetAllItems)
		itemRoutes.GET("/:id", itemController.GetItemByID)
		itemRoutes.PUT("/:id", itemController.UpdateItem)
		itemRoutes.DELETE("/:id", itemController.DeleteItem)
	}

	return router
}
