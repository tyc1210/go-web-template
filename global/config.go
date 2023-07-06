package global

import (
	"github.com/jinzhu/gorm"
	"go-web-template/internal/setting"
)

type Config struct {
	Server   *setting.ServerProperties
	DataBase *setting.DataBaseProperties
	JWT      *setting.JWTProperties
	App      *setting.AppProperties
}

var (
	Cfg *Config
	DB  *gorm.DB
)
