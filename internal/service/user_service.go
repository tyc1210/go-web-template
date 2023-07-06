package service

import (
	"errors"
	"go-web-template/internal/entry/request"
	"go-web-template/internal/entry/response"
	"go-web-template/internal/model"
	"go-web-template/pkg/util"
)

func (service Service) GetUserList(query request.UserQuery) (*response.PageData, error) {
	return service.Dao.GetUserList(query)
}

func (service Service) SaveUser(name, psw string) error {
	dao := service.Dao
	u, err := dao.GetUserByName(name)
	if err != nil {
		return err
	}
	if u != nil {
		return errors.New("用户名已存在")
	}
	return dao.SaveUser(name, psw)
}

func (service Service) Login(name, psw string) (string, error) {
	dao := service.Dao
	u, err := dao.GetUserByName(name)
	if err != nil {
		return "", err
	}
	if u == nil {
		return "", errors.New("用户名或密码错误")
	}
	if psw != u.Password {
		return "", errors.New("用户名或密码错误")
	}
	return util.CreateToken(*u), nil
}

func (service Service) GetById(id uint) (*model.User, error) {
	dao := service.Dao
	return dao.GetUserById(id)
}
