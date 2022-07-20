package router

import (
	"gin-skeleton/app/controller"
	"gin-skeleton/app/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	s := gin.Default()
	s.Use(middleware.CorsMiddleware())
	rootPath := s.Group("/")
	{
		rootPath.GET("/login", controller.NewLoginController().Login)
		rootPath.Any("/logout", controller.NewLoginController().Logout)
	}
	userRoute := s.Group("/user")
	{
		userRoute.GET("/info", controller.Info)
	}
	return s
}
