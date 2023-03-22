package config

import (
	"github.com/gin-gonic/gin"
	"proj1/pkg/controller"
	"proj1/pkg/middleware"
)

func RouterConfig(router *gin.Engine) {
	router.Use(middleware.JWTAuth())
	userController := controller.NewUserController()

	router.POST("/login", userController.Login)
	router.POST("/logout", userController.Logout)
	router.POST("/users", userController.GetAllUsers)
}
