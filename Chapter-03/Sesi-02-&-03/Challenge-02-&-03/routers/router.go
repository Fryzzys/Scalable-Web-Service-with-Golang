package routers

import (
	"simple-rest-api-with-jwt/database"
	"simple-rest-api-with-jwt/controllers"
	"simple-rest-api-with-jwt/middleware"
	"simple-rest-api-with-jwt/repository"
	"simple-rest-api-with-jwt/services"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db := database.GetDB()

	userRepo := repository.UserRepoImpl(db)
	userSvc := services.UserSvcImpl(userRepo)
	userHdl := controllers.UserHdlImpl(userSvc)

	productRepo := repository.ProductRepoImpl(db)
	productSvc := services.ProductSvcImpl(productRepo)
	productHdl := controllers.ProductHdlImpl(productSvc)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		userRouter := v1.Group("/users")
		{
			userRouter.POST("/register", userHdl.Register)
			userRouter.POST("/login", userHdl.Login)
		}

		productRouter := v1.Group("/products")
		{
			productRouter.Use(middleware.Authentication())
			productRouter.POST("/", productHdl.CreateHdl)

			productAuthorizedRouter := productRouter.Group("/")
			{
				productAuthorizedRouter.Use(middleware.ProductAuthorization())
				productAuthorizedRouter.GET("/", productHdl.FindAllHdl)
				productAuthorizedRouter.GET("/:productId", productHdl.FindByIdHdl)
				productAuthorizedRouter.PUT("/:productId", productHdl.UpdateHdl)
				productAuthorizedRouter.DELETE("/:productId", productHdl.DeleteHdl)
			}
		}
	}

	return r
}