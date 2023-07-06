package app

import (
	"github.com/gin-gonic/gin"
	"go-web-template/pkg/errcode"
	"net/http"
)

type CommonResult struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewCommonResult(ctx *gin.Context) *CommonResult {
	return &CommonResult{ctx}
}

func (result CommonResult) Success(data interface{}) {
	result.Ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.GetCode(),
		"data": data,
		"msg":  errcode.Success.GetMsg(),
	})
}

func (result CommonResult) Error(c *errcode.Code) {
	result.Ctx.JSON(http.StatusOK, gin.H{
		"code": c.GetCode(),
		"data": nil,
		"msg":  c.GetMsg(),
	})
}
