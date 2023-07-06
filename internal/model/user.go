package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model        // 嵌入自定义gorm.Model的字段
	Username   string `gorm:"column:username;uniqueIndex;type:varchar(50)" json:"username" form:"username" binding:"required"`
	Password   string `gorm:"column:password;type:varchar(32)" json:"password" form:"password" binding:"required"`
}

func (User) TableName() string {
	return "sys_user"
}
