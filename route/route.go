package route

import (
	"github.com/gin-gonic/gin"
	"github.com/haviz000/API-multi-level-user-golang/controllers/controller_impl"
	"github.com/haviz000/API-multi-level-user-golang/middleware"
	"github.com/haviz000/API-multi-level-user-golang/repository/repository_impl"
	"github.com/haviz000/API-multi-level-user-golang/service/service_impl"
	"gorm.io/gorm"
)

func Route(router *gin.Engine, db *gorm.DB) {
	userRepository := repository_impl.NewUserRepository(db)
	productRepository := repository_impl.NewProductRepository(db)

	userService := service_impl.NewUserService(userRepository)
	productService := service_impl.NewProductService(productRepository, userRepository)

	userController := controller_impl.NewUserController(userService)
	productController := controller_impl.NewProductController(productService)

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", userController.Registration)
		userRouter.POST("/login", userController.Login)
	}

	productRouter := router.Group("/product")
	{
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/", productController.GetProductByRole)
		productRouter.GET(":product_id", productController.GetOneProduct)
		productRouter.PUT(":product_id", productController.UpdateProduct)
		adminRouter := productRouter.Group("/", middleware.AdminMiddleware)
		{
			adminRouter.GET("/all", productController.GetAllProduct)
			adminRouter.DELETE(":product_id", productController.DeleteProduct)
		}
	}
}
