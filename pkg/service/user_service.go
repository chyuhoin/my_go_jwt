package service

import (
	"proj1/pkg/dao"
	"proj1/pkg/entity"
)

type UserService struct {
	userMapper *dao.UserMapper
}

func NewUserService() *UserService {
	return &UserService{dao.NewUserMapper()}
}

func (service *UserService) Login(user *entity.User) bool {
	return service.userMapper.CheckUserWithNameAndPassword(user.Username, user.Password)
}
