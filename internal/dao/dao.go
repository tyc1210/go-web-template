package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Dao {
	return &Dao{DB: db}
}
