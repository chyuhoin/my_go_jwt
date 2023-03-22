package service

import (
	"errors"
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
	if res == nil {
		return nil, errors.New("No ")
	}
	token, err := utils.GenerateToken(res, 60*24)
	if err != nil {
		return nil, err
	}
	return token, nil
}
