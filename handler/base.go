package handler

import (
	"anla.io/taizhou-fe-api/middleware"
	"anla.io/taizhou-fe-api/models"
	"github.com/kataras/iris"
)

// Controller is base controller
type Controller struct{}

// GetUser 获取用户信息
func (ctl Controller) GetUser(ctx iris.Context) models.User {
	return middleware.User{}.GetUser(ctx)
}
