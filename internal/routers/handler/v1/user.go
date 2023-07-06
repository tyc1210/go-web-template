package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-template/internal/entry/request"
	"go-web-template/internal/model"
	"go-web-template/internal/service"
	"go-web-template/pkg/app"
	"go-web-template/pkg/errcode"
	"go-web-template/pkg/util"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// List
// @Summary 获取列表
// @Tags 用户管理
// @Produce  json
// @Param token header string true "token"
// @param data query request.UserQuery true "用户名密码"
// @Success 200 {object} app.Response "成功"
// @Router /api/v1/users [get]
func (u User) List(c *gin.Context) {
	result := app.NewCommonResult(c)
	var query request.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		fmt.Println(err.Error())
		code := errcode.InvalidParams
		code.SetMsg(err.Error())
		result.Error(&code)
		return
	}
	service := service.New(c.Request.Context())
	data, err := service.GetUserList(query)
	if err != nil {
		result.Error(&errcode.ServerError)
		return
	}
	result.Success(data)
}

// GetById
// @Summary 获取单个
// @Tags 用户管理
// @Produce  json
// @Param token header string true "token"
// @Param id path uint true "用户id"
// @Success 200 {object} app.Response "成功"
// @Router /api/v1/users/{id} [get]
func (u User) GetById(c *gin.Context) {
	result := app.NewCommonResult(c)
	service := service.New(c.Request.Context())
	id, err := util.StrToUInt32(c.Param("id"))
	if err != nil {
		result.Error(&errcode.InvalidParams)
	}
	data, err := service.GetById(id)
	if err != nil {
		result.Error(&errcode.ServerError)
		return
	}
	result.Success(data)
}

// Register
// @Summary 注册
// @Tags 用户管理
// @Produce  json
// @Accept json
// @param data body request.LoginRequest true "用户名密码"
// @Success 200 {object} app.Response "成功"
// @Router /api/register [post]
func (u User) Register(c *gin.Context) {
	result := app.NewCommonResult(c)
	service := service.New(c.Request.Context())
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		code := errcode.NewCode(errcode.InvalidParams.GetCode(), err.Error())
		result.Error(&code)
		return
	}
	if err := service.SaveUser(user.Username, user.Password); err != nil {
		code := errcode.NewCode(errcode.UserSave.GetCode(), err.Error())
		result.Error(&code)
		return
	}
	result.Success(nil)
}

// Login
// @Summary 登录
// @Tags 用户管理
// @Produce  json
// @param data body request.LoginRequest true "用户名密码"
// @Success 200 {object} app.Response "成功"
// @Router /api/login [post]
func (u User) Login(c *gin.Context) {
	result := app.NewCommonResult(c)
	service := service.New(c.Request.Context())
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		code := errcode.NewCode(errcode.InvalidParams.GetCode(), err.Error())
		result.Error(&code)
		return
	}
	token, err := service.Login(user.Username, user.Password)
	if err != nil {
		code := errcode.NewCode(errcode.UserLogin.GetCode(), err.Error())
		result.Error(&code)
		return
	}
	result.Success(token)
}
