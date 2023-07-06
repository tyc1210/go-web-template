package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web-template/pkg/app"
	"go-web-template/pkg/errcode"
	"go-web-template/pkg/util"
)

// JWTHandler 解析token并放入context
func JWTHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code = errcode.Success
		token := context.Request.Header.Get("token")
		if token == "" {
			code = errcode.UnauthorizedToken
		}
		user, err := util.ParseToken(token)
		if err != nil {
			code = errcode.UnauthorizedToken
		}

		if code == errcode.Success {
			context.Set("userName", user.Username)
			context.Set("userId", user.ID)
			context.Next()
		} else {
			result := app.NewCommonResult(context)
			result.Error(&code)
			context.Abort()
			return
		}
	}
}
