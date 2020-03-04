package routes

import (
	"github.com/davidchandra95/go-laundry/config"
	"github.com/davidchandra95/go-laundry/modules/user/delivery/http"
	"github.com/davidchandra95/go-laundry/modules/user/repository"
	"github.com/davidchandra95/go-laundry/modules/user/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	db := config.GetDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := http.NewUserHandlers(userService)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			userGroup.GET("/", userHandler.FetchUsers)
			userGroup.GET("/:id", userHandler.GetUser)
		}
	}

	return router
}
