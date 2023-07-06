package service

import (
	"context"
	"go-web-template/global"
	"go-web-template/internal/dao"
)

type Service struct {
	context context.Context
	Dao     *dao.Dao
}

func New(ctx context.Context) Service {
	var service Service
	service.context = ctx
	service.Dao = dao.New(global.DB)
	return service
}
