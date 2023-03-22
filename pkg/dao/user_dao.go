package dao

import (
	"fmt"
	"proj1/pkg/entity"
	"proj1/pkg/utils"
	"xorm.io/xorm"
)

type UserMapper struct {
	engine *xorm.Engine
}

func NewUserMapper() *UserMapper {
	return &UserMapper{engine: utils.GetOrmEngine()}
}

func (mapper *UserMapper) GetUserById(id string) *entity.User {
	var user = entity.User{}
	mapper.engine.Where("id = ?", id).Get(&user)
	return &user
}

func (mapper *UserMapper) GetUserByNameAndPassword(username string, password string) *entity.User {
	var user = entity.User{}
	mapper.engine.Where("username = ?", username).And("password = ?", password).Get(&user)
	return &user
}

func (mapper *UserMapper) CheckUserWithNameAndPassword(username string, password string) bool {
	has, err := mapper.engine.Where("username = ?", username).And("password = ?", password).Exist(&entity.User{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return has
}
