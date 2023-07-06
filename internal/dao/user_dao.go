package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go-web-template/internal/entry/request"
	"go-web-template/internal/entry/response"
	"go-web-template/internal/model"
)

func (d Dao) GetUserList(query request.UserQuery) (*response.PageData, error) {
	db := d.DB
	users := make([]*model.User, query.Size)
	if query.UserName != "" {
		db.Where("username = ?", query.UserName)
	}
	if err := db.Order("created_at desc").Limit(query.Size).Offset(query.Page - 1).Find(&users).Error; err != nil {
		return nil, err
	}
	var total int64 = 0
	db.Model(model.User{}).Count(&total)
	pageData := response.NewPageData(total, users)
	return &pageData, nil
}

func (d Dao) SaveUser(name string, psw string) error {
	u := model.User{Username: name, Password: psw}
	if err := d.DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (d Dao) GetUserByName(name string) (*model.User, error) {
	user := model.User{}
	err := d.DB.Debug().Where("username = ?", name).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (d Dao) GetUserById(id uint) (*model.User, error) {
	user := model.User{}
	err := d.DB.Debug().Where("id = ?", id).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
