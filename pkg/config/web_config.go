package config

import (
	"github.com/gin-gonic/gin"
	"proj1/pkg/controller"
)

func RouterConfig(router *gin.Engine) {
	userController := controller.NewUserController()

	router.POST("/login", userController.Login)
	router.POST("/logout", userController.Logout)
	router.POST("/users", userController.GetAllUsers)
}
