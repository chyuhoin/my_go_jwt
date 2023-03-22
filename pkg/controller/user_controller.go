package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"proj1/pkg/entity"
	"proj1/pkg/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	controller := UserController{service.NewUserService()}
	return &controller
}

func (ctl *UserController) Login(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if ctl.userService.Login(&user) {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "No such people"})
	}
}

func (ctl *UserController) Logout(c *gin.Context) {

}

func (ctl *UserController) GetAllUsers(c *gin.Context) {

}
