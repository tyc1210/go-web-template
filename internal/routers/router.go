package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-web-template/docs"
	"go-web-template/global"
	"go-web-template/internal/middleware"
	v1 "go-web-template/internal/routers/handler/v1"
	"go-web-template/pkg/logger"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())
	router.Use(logger.NewLogger())
	r := router.Group("/api")
	// 文件上传相关
	router.MaxMultipartMemory = global.Cfg.App.UploadMaxSize << 20
	file := v1.NewFile()
	r.POST("/upload", file.Upload2Local)
	r.StaticFS("/static", http.Dir(global.Cfg.App.UploadSavePath))
	// 接口文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := v1.NewUser()

	r.POST("/register", user.Register)
	r.POST("/login", user.Login)

	routerV1 := r.Group("/v1")
	routerV1.Use(middleware.JWTHandler())
	{
		routerV1.GET("/users", user.List)
		routerV1.GET("/users/:id", user.GetById)
	}

	return router
}
