package router

import (
	"github.com/gin-gonic/gin"
	"practice/app/controller"
	"practice/app/middleware"
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
