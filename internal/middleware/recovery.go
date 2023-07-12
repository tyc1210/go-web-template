package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/pkg/app"
	"go-web-template/pkg/errcode"
	"go-web-template/pkg/logger"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("panic recover err: %v", err)
				// todo 进行报警或其它处理
				result := app.NewCommonResult(c)
				result.Error(&errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
