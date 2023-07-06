package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-web-template/global"
	"go-web-template/internal/database"
	"go-web-template/internal/routers"
	"go-web-template/pkg/logger"
	"log"
	"net/http"
	"strings"
)

// @title 模板项目
// @version 1.0
// @description Go 语言开发 web 项目
// @termsOfService https://github.com
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:    ":" + global.Cfg.Server.HttpPort,
		Handler: router,
	}
	s.ListenAndServe()
}

func init() {
	InitConfig()
	InitLogger()
	InitDb()
}

func InitLogger() {
	logger.InitLogger(global.Cfg.App)
}

func InitDb() {
	var err error
	global.DB, err = database.NewDBEngine(global.Cfg.DataBase)
	if err != nil {
		log.Fatalf("InitDb err: %v", err)
	}
}

func InitConfig() {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	viper.AutomaticEnv() // 启用自动读取环境变量的功能
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig() // 找到并加载配置文件
	if err != nil {             // 处理错误
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.Unmarshal(&global.Cfg); err != nil {
		log.Printf("unmarshal config file failed, %v", err)
	}
	log.Printf("server http port ===================> %s", global.Cfg.Server.HttpPort)
}
