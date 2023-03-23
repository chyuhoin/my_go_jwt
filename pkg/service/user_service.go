package service

import (
	"fmt"
	"proj1/pkg/dao"
	"proj1/pkg/entity"
	"proj1/pkg/utils"
)

type UserService struct {
	userMapper *dao.UserMapper
}

func NewUserService() *UserService {
	return &UserService{dao.NewUserMapper()}
}

func (service *UserService) Login(user *entity.User) (*string, error) {
	res := service.userMapper.GetUserByNameAndPassword(user.Username, user.Password)
	fmt.Println(res.Id)
	if res.Id == "" {
		return nil, utils.NoSuchUser
	}
	token, err := utils.GenerateToken(res, 60*24)
	if err != nil {
		return nil, err
	}
	user.Id = res.Id
	user.Role = res.Role
	return token, nil
}
