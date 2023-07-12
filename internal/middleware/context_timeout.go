package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前请求上下问，并根据时间创建一个带有超时时间的上下文
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		// 取消函数 用于在需要时取消上下文
		defer cancel()
		// 将新创建的上下文ctx设置为当前请求的上下文
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
